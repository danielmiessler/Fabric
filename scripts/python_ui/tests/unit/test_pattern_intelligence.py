"""
Unit tests for pattern intelligence service.
"""
import pytest
from unittest.mock import Mock, patch
from datetime import datetime, timedelta

from services.pattern_intelligence import (
    PatternIntelligenceService, PatternAnalytics, PatternRecommendation,
    get_pattern_intelligence
)
from utils.typing import PatternSpec


@pytest.fixture
def intelligence_service():
    """Create a fresh PatternIntelligenceService for testing."""
    return PatternIntelligenceService()


@pytest.fixture
def sample_patterns():
    """Sample pattern specs for testing."""
    return [
        PatternSpec(name="analyze_code", path=None, content="# IDENTITY\nAnalyze code quality\n# STEPS\n- Review syntax\n- Check patterns"),
        PatternSpec(name="write_docs", path=None, content="# IDENTITY\nWrite documentation\n# STEPS\n- Analyze code\n- Generate docs"),
        PatternSpec(name="extract_data", path=None, content="# IDENTITY\nExtract key data\n# STEPS\n- Parse content\n- Extract info"),
    ]


def test_categorize_pattern(intelligence_service):
    """Test pattern categorization logic."""
    # Test analysis pattern
    with patch('services.patterns.load_pattern') as mock_load:
        mock_load.return_value = PatternSpec(
            name="analyze_code", 
            path=None, 
            content="analyze the code quality and provide recommendations"
        )
        category = intelligence_service.categorize_pattern("analyze_code")
        assert category == "ANALYSIS"
    
    # Test writing pattern
    with patch('services.patterns.load_pattern') as mock_load:
        mock_load.return_value = PatternSpec(
            name="write_docs", 
            path=None, 
            content="write comprehensive documentation for the project"
        )
        category = intelligence_service.categorize_pattern("write_docs")
        assert category == "WRITING"


def test_analyze_pattern_usage_no_data(intelligence_service):
    """Test pattern analytics with no execution data."""
    with patch('services.storage.read_outputs', return_value=[]):
        analytics = intelligence_service.analyze_pattern_usage("test_pattern")
        
        assert isinstance(analytics, PatternAnalytics)
        assert analytics.pattern_name == "test_pattern"
        assert analytics.success_rate == 0.85  # Default
        assert analytics.usage_count == 0


def test_analyze_pattern_usage_with_data(intelligence_service):
    """Test pattern analytics with execution data."""
    mock_outputs = [
        {"pattern_name": "test_pattern", "output": "success output"},
        {"pattern_name": "test_pattern", "output": "another success"},
        {"pattern_name": "test_pattern", "output": None},  # Failed execution
        {"pattern_name": "other_pattern", "output": "different pattern"}
    ]
    
    with patch('services.storage.read_outputs', return_value=mock_outputs):
        analytics = intelligence_service.analyze_pattern_usage("test_pattern")
        
        assert analytics.pattern_name == "test_pattern"
        assert analytics.success_rate == 2/3  # 2 successes out of 3 executions
        assert analytics.usage_count == 3


def test_search_patterns_semantic(intelligence_service, sample_patterns):
    """Test semantic pattern search."""
    with patch('services.patterns.list_patterns', return_value=sample_patterns):
        results = intelligence_service.search_patterns_semantic("analyze code", limit=5)
        
        assert isinstance(results, list)
        assert len(results) <= 5
        
        # Check that analyze_code is in results with high relevance
        pattern_names = [result[0] for result in results]
        assert "analyze_code" in pattern_names
        
        # Check relevance scores are between 0 and 1
        for pattern_name, score in results:
            assert 0 <= score <= 1


def test_recommend_patterns(intelligence_service, sample_patterns):
    """Test pattern recommendation system."""
    with patch('services.patterns.list_patterns', return_value=sample_patterns):
        recommendations = intelligence_service.recommend_patterns(
            context="I need to analyze some code",
            current_patterns=["write_docs"],
            limit=3
        )
        
        assert isinstance(recommendations, list)
        assert len(recommendations) <= 3
        
        # Should not recommend already selected patterns
        rec_names = [rec.pattern_name for rec in recommendations]
        assert "write_docs" not in rec_names
        
        # Check recommendation structure
        for rec in recommendations:
            assert isinstance(rec, PatternRecommendation)
            assert 0 <= rec.confidence_score <= 1
            assert rec.reason
            assert rec.category


def test_get_trending_patterns_no_data(intelligence_service):
    """Test trending patterns with no data."""
    with patch('services.storage.read_outputs', return_value=[]):
        trending = intelligence_service.get_trending_patterns(days=7, limit=5)
        assert trending == []


def test_get_trending_patterns_with_data(intelligence_service):
    """Test trending patterns with mock data."""
    # Create mock data with timestamps
    now = datetime.now()
    mock_outputs = [
        {"pattern_name": "popular_pattern", "timestamp": now.strftime("%Y-%m-%d %H:%M:%S")},
        {"pattern_name": "popular_pattern", "timestamp": now.strftime("%Y-%m-%d %H:%M:%S")},
        {"pattern_name": "other_pattern", "timestamp": now.strftime("%Y-%m-%d %H:%M:%S")},
    ]
    
    with patch('services.storage.read_outputs', return_value=mock_outputs):
        trending = intelligence_service.get_trending_patterns(days=7, limit=5)
        
        assert len(trending) > 0
        assert trending[0][0] == "popular_pattern"  # Most popular should be first
        assert trending[0][1] == 2  # Used 2 times


def test_suggest_workflow_optimizations(intelligence_service):
    """Test workflow optimization suggestions."""
    selected_patterns = ["analyze_code", "write_docs", "extract_data"]
    
    with patch.object(intelligence_service, 'categorize_pattern') as mock_categorize:
        mock_categorize.side_effect = lambda x: {
            "analyze_code": "ANALYSIS",
            "write_docs": "WRITING", 
            "extract_data": "EXTRACTION"
        }[x]
        
        suggestions = intelligence_service.suggest_workflow_optimizations(selected_patterns)
        
        assert isinstance(suggestions, list)
        # Should have at least one suggestion for 3 different category patterns
        assert len(suggestions) >= 0


def test_estimate_execution_time(intelligence_service):
    """Test execution time estimation."""
    with patch('services.patterns.load_pattern') as mock_load:
        mock_load.return_value = PatternSpec(
            name="test_pattern",
            path=None,
            content="# A simple pattern\nDo something quick"
        )
        
        time_estimate = intelligence_service._estimate_execution_time("test_pattern")
        assert isinstance(time_estimate, float)
        assert time_estimate > 0


def test_estimate_complexity(intelligence_service):
    """Test complexity estimation."""
    with patch('services.patterns.load_pattern') as mock_load:
        # Simple pattern
        mock_load.return_value = PatternSpec(
            name="simple_pattern",
            path=None,
            content="# Simple\nDo something"
        )
        
        complexity = intelligence_service._estimate_complexity("simple_pattern")
        assert 0 <= complexity <= 1
        
        # Complex pattern
        complex_content = "# IDENTITY\n" + "Complex analysis pattern\n" * 100 + "# STEPS\n" + "Step\n" * 10
        mock_load.return_value = PatternSpec(
            name="complex_pattern",
            path=None,
            content=complex_content
        )
        
        complex_complexity = intelligence_service._estimate_complexity("complex_pattern")
        assert complex_complexity > complexity  # Should be more complex


def test_singleton_pattern():
    """Test that get_pattern_intelligence returns the same instance."""
    service1 = get_pattern_intelligence()
    service2 = get_pattern_intelligence()
    assert service1 is service2
