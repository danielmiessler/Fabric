"""
Pattern intelligence service providing AI-powered pattern analysis, recommendations, and insights.
"""
from __future__ import annotations
import re
import json
from dataclasses import dataclass, field
from datetime import datetime, timedelta
from typing import List, Dict, Any, Optional, Tuple
from pathlib import Path

from utils.logging import logger
from utils.typing import PatternSpec
from services import patterns, storage
from services.dependencies import OptionalImport, check_and_install_if_missing


@dataclass
class PatternAnalytics:
    """Analytics data for a pattern."""
    pattern_name: str
    success_rate: float = 0.0
    avg_execution_time: float = 0.0
    complexity_score: float = 0.5
    usage_count: int = 0
    last_used: Optional[datetime] = None
    common_inputs: List[str] = field(default_factory=list)
    failure_reasons: List[str] = field(default_factory=list)


@dataclass 
class PatternRecommendation:
    """AI-powered pattern recommendation."""
    pattern_name: str
    confidence_score: float
    reason: str
    category: str
    estimated_execution_time: float = 0.0
    similarity_score: float = 0.0


@dataclass
class PatternRelationship:
    """Relationship between two patterns."""
    pattern_a: str
    pattern_b: str
    relationship_type: str  # "sequential", "complementary", "alternative"
    confidence: float
    description: str


class PatternIntelligenceService:
    """Service providing AI-powered pattern analysis and recommendations."""
    
    def __init__(self):
        # Use best available cache implementation
        from utils.dependency_fallbacks import get_cache_implementation
        
        self._analytics_cache = get_cache_implementation(maxsize=256, ttl=600)  # 10 min TTL
        self._category_cache = get_cache_implementation(maxsize=512, ttl=3600)  # 1 hour TTL
        self._relationships_cache = get_cache_implementation(maxsize=128, ttl=1800)  # 30 min TTL
        self._cache_expiry = timedelta(minutes=10)
        self._last_cache_refresh = datetime.now()
        
        # Pattern categories based on common Fabric patterns
        self._categories = {
            "ANALYSIS": ["analyze", "summarize", "extract", "review", "audit", "inspect"],
            "WRITING": ["write", "create", "generate", "compose", "draft", "author"],
            "CODING": ["code", "program", "debug", "refactor", "test", "implement"],
            "COMMUNICATION": ["email", "message", "letter", "proposal", "report"],
            "EDUCATION": ["explain", "teach", "learn", "tutorial", "lesson", "guide"],
            "PLANNING": ["plan", "strategy", "roadmap", "schedule", "organize"],
            "TRANSFORMATION": ["convert", "transform", "translate", "format", "restructure"],
            "EXTRACTION": ["extract", "parse", "scrape", "mine", "collect", "gather"],
            "SUMMARY": ["summarize", "brief", "outline", "digest", "condensed"],
            "CREATIVITY": ["creative", "brainstorm", "ideate", "imagine", "design"]
        }

    def search_patterns_semantic(
        self, 
        query: str, 
        limit: int = 10, 
        filters: Optional[Dict[str, Any]] = None
    ) -> List[Tuple[str, float]]:
        """
        Search patterns using semantic similarity and keyword matching.
        
        Args:
            query: Search query
            limit: Maximum number of results
            filters: Optional filters (max_complexity, max_execution_time, category)
            
        Returns:
            List of (pattern_name, relevance_score) tuples
        """
        try:
            pattern_specs = patterns.list_patterns()
            if not pattern_specs:
                return []
            
            results = []
            query_lower = query.lower().strip()
            
            for spec in pattern_specs:
                relevance_score = self._calculate_relevance_score(spec, query_lower)
                
                # Apply filters
                if filters and not self._passes_filters(spec, filters):
                    continue
                
                if relevance_score > 0.1:  # Minimum relevance threshold
                    results.append((spec.name, relevance_score))
            
            # Sort by relevance score descending
            results.sort(key=lambda x: x[1], reverse=True)
            return results[:limit]
            
        except Exception as e:
            logger.error(f"Semantic search failed: {e}")
            return []
    
    def recommend_patterns(
        self,
        context: str = "",
        current_patterns: List[str] = None,
        limit: int = 5
    ) -> List[PatternRecommendation]:
        """
        Generate AI-powered pattern recommendations based on context and current selection.
        
        Args:
            context: Input context for recommendations
            current_patterns: Currently selected patterns
            limit: Maximum number of recommendations
            
        Returns:
            List of pattern recommendations
        """
        try:
            if current_patterns is None:
                current_patterns = []
            
            pattern_specs = patterns.list_patterns()
            recommendations = []
            
            for spec in pattern_specs:
                if spec.name in current_patterns:
                    continue  # Don't recommend already selected patterns
                
                confidence = self._calculate_recommendation_confidence(
                    spec, context, current_patterns
                )
                
                if confidence > 0.3:  # Minimum confidence threshold
                    category = self.categorize_pattern(spec.name)
                    analytics = self.analyze_pattern_usage(spec.name)
                    
                    reason = self._generate_recommendation_reason(
                        spec, context, current_patterns, confidence
                    )
                    
                    recommendation = PatternRecommendation(
                        pattern_name=spec.name,
                        confidence_score=confidence,
                        reason=reason,
                        category=category,
                        estimated_execution_time=analytics.avg_execution_time
                    )
                    recommendations.append(recommendation)
            
            # Sort by confidence
            recommendations.sort(key=lambda x: x.confidence_score, reverse=True)
            return recommendations[:limit]
            
        except Exception as e:
            logger.error(f"Pattern recommendation failed: {e}")
            return []
    
    def categorize_pattern(self, pattern_name: str) -> str:
        """
        Categorize a pattern based on its name and content.
        
        Args:
            pattern_name: Name of the pattern
            
        Returns:
            Category name
        """
        # Check cache first
        if pattern_name in self._category_cache:
            return self._category_cache[pattern_name]
        
        try:
            # Load pattern content
            pattern_spec = patterns.load_pattern(pattern_name)
            content = (pattern_spec.content or "").lower()
            name_lower = pattern_name.lower()
            
            # Score each category
            category_scores = {}
            for category, keywords in self._categories.items():
                score = 0
                for keyword in keywords:
                    # Check in pattern name (weighted higher)
                    if keyword in name_lower:
                        score += 2
                    # Check in content
                    if keyword in content:
                        score += 1
                
                category_scores[category] = score
            
            # Get highest scoring category
            best_category = max(category_scores, key=category_scores.get)
            if category_scores[best_category] == 0:
                best_category = "GENERAL"
            
            # Cache the result
            self._category_cache[pattern_name] = best_category
            return best_category
            
        except Exception as e:
            logger.warning(f"Pattern categorization failed for {pattern_name}: {e}")
            return "GENERAL"
    
    def analyze_pattern_usage(self, pattern_name: str) -> PatternAnalytics:
        """
        Analyze pattern usage statistics and performance.
        
        Args:
            pattern_name: Name of the pattern to analyze
            
        Returns:
            Pattern analytics data
        """
        # Check cache
        if (pattern_name in self._analytics_cache and 
            datetime.now() - self._last_cache_refresh < self._cache_expiry):
            return self._analytics_cache[pattern_name]
        
        try:
            # Get execution data from storage and session
            outputs = storage.read_outputs()
            session_outputs = []
            try:
                import streamlit as st
                session_outputs = st.session_state.get("output_logs", [])
            except:
                pass
            
            # Combine all execution data
            all_outputs = session_outputs + outputs
            
            # Filter for this pattern
            pattern_executions = [
                log for log in all_outputs 
                if log.get("pattern_name") == pattern_name or log.get("pattern") == pattern_name
            ]
            
            if not pattern_executions:
                # Return default analytics for patterns with no execution history
                analytics = PatternAnalytics(
                    pattern_name=pattern_name,
                    success_rate=0.85,  # Default assumption
                    avg_execution_time=self._estimate_execution_time(pattern_name),
                    complexity_score=self._estimate_complexity(pattern_name),
                    usage_count=0
                )
            else:
                # Calculate real analytics
                successful_runs = len([log for log in pattern_executions if log.get("output")])
                success_rate = successful_runs / len(pattern_executions) if pattern_executions else 0.0
                
                # Estimate execution time (we don't track this yet, so use heuristic)
                avg_time = self._estimate_execution_time(pattern_name)
                
                analytics = PatternAnalytics(
                    pattern_name=pattern_name,
                    success_rate=success_rate,
                    avg_execution_time=avg_time,
                    complexity_score=self._estimate_complexity(pattern_name),
                    usage_count=len(pattern_executions),
                    last_used=datetime.now() if pattern_executions else None
                )
            
            # Cache the result
            self._analytics_cache[pattern_name] = analytics
            return analytics
            
        except Exception as e:
            logger.error(f"Pattern analysis failed for {pattern_name}: {e}")
            # Return default analytics
            return PatternAnalytics(
                pattern_name=pattern_name,
                success_rate=0.75,
                avg_execution_time=30.0,
                complexity_score=0.5
            )
    
    def get_trending_patterns(self, days: int = 7, limit: int = 10) -> List[Tuple[str, int]]:
        """
        Get trending patterns based on recent usage.
        
        Args:
            days: Number of days to look back
            limit: Maximum number of results
            
        Returns:
            List of (pattern_name, usage_count) tuples
        """
        try:
            # Get recent executions
            cutoff_date = datetime.now() - timedelta(days=days)
            outputs = storage.read_outputs()
            
            # Filter for recent executions
            recent_outputs = []
            for log in outputs:
                try:
                    # Parse timestamp
                    if "timestamp" in log:
                        timestamp_str = log["timestamp"]
                        # Handle different timestamp formats
                        for fmt in ["%Y-%m-%d %H:%M:%S", "%Y-%m-%dT%H:%M:%SZ", "%Y-%m-%dT%H:%M:%S"]:
                            try:
                                timestamp = datetime.strptime(timestamp_str, fmt)
                                if timestamp >= cutoff_date:
                                    recent_outputs.append(log)
                                break
                            except ValueError:
                                continue
                except Exception:
                    continue
            
            # Count pattern usage
            pattern_counts = {}
            for log in recent_outputs:
                pattern = log.get("pattern_name") or log.get("pattern", "Unknown")
                pattern_counts[pattern] = pattern_counts.get(pattern, 0) + 1
            
            # Sort by usage count
            trending = sorted(pattern_counts.items(), key=lambda x: x[1], reverse=True)
            return trending[:limit]
            
        except Exception as e:
            logger.error(f"Trending patterns analysis failed: {e}")
            return []
    
    def get_pattern_relationships(self, pattern_name: str) -> List[PatternRelationship]:
        """
        Find relationships between patterns.
        
        Args:
            pattern_name: Name of the pattern
            
        Returns:
            List of pattern relationships
        """
        # Check cache
        if pattern_name in self._relationships_cache:
            return self._relationships_cache[pattern_name]
        
        try:
            relationships = []
            pattern_specs = patterns.list_patterns()
            target_spec = patterns.load_pattern(pattern_name)
            
            for spec in pattern_specs:
                if spec.name == pattern_name:
                    continue
                
                # Calculate relationship
                relationship = self._analyze_pattern_relationship(target_spec, spec)
                if relationship:
                    relationships.append(relationship)
            
            # Cache the result
            self._relationships_cache[pattern_name] = relationships
            return relationships
            
        except Exception as e:
            logger.error(f"Pattern relationship analysis failed for {pattern_name}: {e}")
            return []
    
    def suggest_workflow_optimizations(self, selected_patterns: List[str]) -> List[Dict[str, Any]]:
        """
        Suggest workflow optimizations for selected patterns.
        
        Args:
            selected_patterns: List of selected pattern names
            
        Returns:
            List of optimization suggestions
        """
        suggestions = []
        
        try:
            if len(selected_patterns) < 2:
                return suggestions
            
            # Analyze pattern categories for grouping opportunities
            categories = {}
            for pattern in selected_patterns:
                category = self.categorize_pattern(pattern)
                if category not in categories:
                    categories[category] = []
                categories[category].append(pattern)
            
            # Suggest parallel execution for similar categories
            for category, patterns_list in categories.items():
                if len(patterns_list) > 1:
                    suggestions.append({
                        "type": "parallel_execution",
                        "title": f"Parallelize {category} patterns",
                        "description": f"Patterns {', '.join(patterns_list)} can run in parallel",
                        "severity": "medium",
                        "parallel_patterns": patterns_list
                    })
            
            # Suggest reordering for efficiency
            if len(selected_patterns) > 2:
                # Sort by estimated execution time (fastest first)
                pattern_times = []
                for pattern in selected_patterns:
                    analytics = self.analyze_pattern_usage(pattern)
                    pattern_times.append((pattern, analytics.avg_execution_time))
                
                pattern_times.sort(key=lambda x: x[1])
                suggested_order = [p[0] for p in pattern_times]
                
                if suggested_order != selected_patterns:
                    suggestions.append({
                        "type": "reorder_patterns",
                        "title": "Optimize execution order",
                        "description": "Run faster patterns first to get quicker feedback",
                        "severity": "low",
                        "suggested_order": suggested_order
                    })
            
            # Suggest consolidation for redundant patterns
            redundant_patterns = self._find_redundant_patterns(selected_patterns)
            if redundant_patterns:
                suggestions.append({
                    "type": "consolidate_patterns",
                    "title": "Consider consolidating similar patterns",
                    "description": f"Patterns {', '.join(redundant_patterns)} have similar functionality",
                    "severity": "low",
                    "redundant_patterns": redundant_patterns
                })
                
        except Exception as e:
            logger.error(f"Workflow optimization failed: {e}")
        
        return suggestions
    
    def _calculate_relevance_score(self, spec: PatternSpec, query: str) -> float:
        """Calculate relevance score for a pattern against a search query."""
        score = 0.0
        name_lower = spec.name.lower()
        content_lower = (spec.content or "").lower()
        
        # Exact name match gets highest score
        if query == name_lower:
            score += 1.0
        elif query in name_lower:
            score += 0.8
        
        # Partial name matches
        query_words = query.split()
        for word in query_words:
            if word in name_lower:
                score += 0.4 / len(query_words)
        
        # Content matches (weighted lower)
        for word in query_words:
            if word in content_lower:
                score += 0.2 / len(query_words)
        
        # Category relevance
        category = self.categorize_pattern(spec.name)
        if any(keyword in query for keyword in self._categories.get(category, [])):
            score += 0.3
        
        return min(score, 1.0)  # Cap at 1.0
    
    def _passes_filters(self, spec: PatternSpec, filters: Dict[str, Any]) -> bool:
        """Check if pattern passes the specified filters."""
        if not filters:
            return True
        
        # Complexity filter
        if "max_complexity" in filters:
            complexity = self._estimate_complexity(spec.name)
            if complexity > filters["max_complexity"]:
                return False
        
        # Execution time filter
        if "max_execution_time" in filters:
            est_time = self._estimate_execution_time(spec.name)
            if est_time > filters["max_execution_time"]:
                return False
        
        # Category filter
        if "category" in filters and filters["category"]:
            category = self.categorize_pattern(spec.name)
            if category not in filters["category"]:
                return False
        
        return True
    
    def _calculate_recommendation_confidence(
        self, 
        spec: PatternSpec, 
        context: str, 
        current_patterns: List[str]
    ) -> float:
        """Calculate confidence score for recommending a pattern."""
        confidence = 0.0
        
        # Context relevance
        if context:
            context_lower = context.lower()
            content_lower = (spec.content or "").lower()
            name_lower = spec.name.lower()
            
            # Look for keyword matches
            context_words = set(context_lower.split())
            content_words = set(content_lower.split())
            name_words = set(name_lower.replace("_", " ").split())
            
            # Calculate word overlap
            context_content_overlap = len(context_words & content_words) / max(len(context_words), 1)
            context_name_overlap = len(context_words & name_words) / max(len(context_words), 1)
            
            confidence += context_content_overlap * 0.4 + context_name_overlap * 0.6
        
        # Category complementarity with current patterns
        if current_patterns:
            current_categories = {self.categorize_pattern(p) for p in current_patterns}
            pattern_category = self.categorize_pattern(spec.name)
            
            # Boost confidence for complementary categories
            complementary_pairs = {
                ("ANALYSIS", "WRITING"), ("EXTRACTION", "SUMMARY"),
                ("PLANNING", "WRITING"), ("CREATIVITY", "ANALYSIS")
            }
            
            for cat_a, cat_b in complementary_pairs:
                if (pattern_category == cat_a and cat_b in current_categories) or \
                   (pattern_category == cat_b and cat_a in current_categories):
                    confidence += 0.3
                    break
        
        # Usage analytics boost
        analytics = self.analyze_pattern_usage(spec.name)
        if analytics.success_rate > 0.8:
            confidence += 0.2
        
        return min(confidence, 1.0)
    
    def _generate_recommendation_reason(
        self, 
        spec: PatternSpec, 
        context: str, 
        current_patterns: List[str], 
        confidence: float
    ) -> str:
        """Generate human-readable reason for recommendation."""
        reasons = []
        
        if confidence > 0.7:
            reasons.append("Highly relevant to your input")
        elif confidence > 0.5:
            reasons.append("Good match for your workflow")
        else:
            reasons.append("May complement your selection")
        
        # Category-based reasons
        category = self.categorize_pattern(spec.name)
        current_categories = {self.categorize_pattern(p) for p in current_patterns}
        
        if category not in current_categories:
            reasons.append(f"Adds {category.lower()} capability")
        
        # Usage-based reasons
        analytics = self.analyze_pattern_usage(spec.name)
        if analytics.success_rate > 0.8:
            reasons.append("High success rate")
        
        return " • ".join(reasons)
    
    def _estimate_execution_time(self, pattern_name: str) -> float:
        """Estimate execution time based on pattern characteristics."""
        try:
            pattern_spec = patterns.load_pattern(pattern_name)
            content_length = len(pattern_spec.content or "")
            
            # Base time estimation
            base_time = 15.0  # Base 15 seconds
            
            # Adjust based on content length and complexity
            complexity_factor = min(content_length / 1000.0, 3.0)  # Max 3x multiplier
            
            # Adjust based on pattern type
            name_lower = pattern_name.lower()
            if any(word in name_lower for word in ["analyze", "summary", "extract"]):
                base_time += 10.0  # Analysis patterns take longer
            elif any(word in name_lower for word in ["write", "create", "generate"]):
                base_time += 20.0  # Creative patterns take longer
            elif any(word in name_lower for word in ["simple", "quick", "basic"]):
                base_time = max(base_time - 5.0, 5.0)  # Simple patterns are faster
            
            return base_time + (complexity_factor * 5.0)
            
        except Exception:
            return 30.0  # Default 30 seconds
    
    def _estimate_complexity(self, pattern_name: str) -> float:
        """Estimate pattern complexity (0.0 to 1.0)."""
        try:
            pattern_spec = patterns.load_pattern(pattern_name)
            content = pattern_spec.content or ""
            
            complexity_score = 0.0
            
            # Length-based complexity
            content_length = len(content)
            if content_length > 2000:
                complexity_score += 0.3
            elif content_length > 1000:
                complexity_score += 0.2
            elif content_length > 500:
                complexity_score += 0.1
            
            # Structure complexity
            sections = len(re.findall(r'^#\s+', content, re.MULTILINE))
            if sections > 5:
                complexity_score += 0.2
            elif sections > 3:
                complexity_score += 0.1
            
            # Variable complexity
            variables = len(re.findall(r'\{\{(\w+)\}\}', content))
            if variables > 3:
                complexity_score += 0.3
            elif variables > 0:
                complexity_score += 0.1
            
            # Keyword-based complexity
            complex_keywords = ["analyze", "extract", "transform", "compare", "evaluate"]
            for keyword in complex_keywords:
                if keyword in pattern_name.lower() or keyword in content.lower():
                    complexity_score += 0.1
                    break
            
            return min(complexity_score, 1.0)
            
        except Exception:
            return 0.5  # Default medium complexity
    
    def _analyze_pattern_relationship(
        self, 
        spec_a: PatternSpec, 
        spec_b: PatternSpec
    ) -> Optional[PatternRelationship]:
        """Analyze relationship between two patterns."""
        try:
            cat_a = self.categorize_pattern(spec_a.name)
            cat_b = self.categorize_pattern(spec_b.name)
            
            # Sequential relationships
            sequential_pairs = [
                ("EXTRACTION", "ANALYSIS"), ("ANALYSIS", "SUMMARY"),
                ("PLANNING", "WRITING"), ("CREATIVITY", "ANALYSIS")
            ]
            
            for first, second in sequential_pairs:
                if cat_a == first and cat_b == second:
                    return PatternRelationship(
                        pattern_a=spec_a.name,
                        pattern_b=spec_b.name,
                        relationship_type="sequential",
                        confidence=0.8,
                        description=f"{first} followed by {second}"
                    )
            
            # Complementary relationships (same category)
            if cat_a == cat_b and cat_a != "GENERAL":
                return PatternRelationship(
                    pattern_a=spec_a.name,
                    pattern_b=spec_b.name,
                    relationship_type="complementary",
                    confidence=0.6,
                    description=f"Both are {cat_a.lower()} patterns"
                )
            
            return None
            
        except Exception as e:
            logger.warning(f"Relationship analysis failed: {e}")
            return None
    
    def _find_redundant_patterns(self, selected_patterns: List[str]) -> List[str]:
        """Find patterns that might be redundant."""
        redundant = []
        
        try:
            for i, pattern_a in enumerate(selected_patterns):
                for pattern_b in selected_patterns[i+1:]:
                    # Check category similarity
                    cat_a = self.categorize_pattern(pattern_a)
                    cat_b = self.categorize_pattern(pattern_b)
                    
                    # Check name similarity
                    name_similarity = self._calculate_name_similarity(pattern_a, pattern_b)
                    
                    if cat_a == cat_b and name_similarity > 0.7:
                        if pattern_a not in redundant:
                            redundant.append(pattern_a)
                        if pattern_b not in redundant:
                            redundant.append(pattern_b)
                            
        except Exception as e:
            logger.error(f"Redundancy analysis failed: {e}")
        
        return redundant
    
    def _calculate_name_similarity(self, name_a: str, name_b: str) -> float:
        """Calculate similarity between two pattern names."""
        # Simple word-based similarity
        words_a = set(name_a.lower().replace("_", " ").split())
        words_b = set(name_b.lower().replace("_", " ").split())
        
        if not words_a or not words_b:
            return 0.0
        
        intersection = len(words_a & words_b)
        union = len(words_a | words_b)
        
        return intersection / union if union > 0 else 0.0


# Singleton instance
_pattern_intelligence_service = None


def get_pattern_intelligence() -> PatternIntelligenceService:
    """Get singleton PatternIntelligenceService instance."""
    global _pattern_intelligence_service
    if _pattern_intelligence_service is None:
        _pattern_intelligence_service = PatternIntelligenceService()
    return _pattern_intelligence_service


# Convenience module-level instance for easy access
pattern_intelligence = get_pattern_intelligence()
