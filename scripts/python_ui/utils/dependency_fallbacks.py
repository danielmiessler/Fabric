"""
Fallback implementations for optional dependencies.
Provides graceful degradation when advanced features are not available.
"""
from typing import Any, Dict, List, Optional
import streamlit as st


class PlotlyFallback:
    """Fallback implementation for plotly functionality."""
    
    @staticmethod
    def create_timeline_chart(data: List[Dict[str, Any]]) -> None:
        """Create a simple timeline visualization without plotly."""
        st.markdown("### 📈 Execution Timeline")
        st.info("📊 Enhanced timeline charts available with plotly installation")
        
        # Simple table fallback
        if data:
            st.dataframe(data, use_container_width=True)
        else:
            st.info("No execution data to display")
    
    @staticmethod
    def create_performance_chart(performance_data: List[Dict[str, Any]]) -> None:
        """Create a simple performance visualization without plotly."""
        st.markdown("### 📊 Performance Trends")
        st.info("📈 Advanced performance charts available with plotly installation")
        
        # Simple metrics fallback
        if performance_data:
            st.dataframe(performance_data, use_container_width=True)
        else:
            st.info("No performance data to display")


class SentenceTransformersFallback:
    """Fallback implementation for semantic search without sentence-transformers."""
    
    @staticmethod
    def encode_patterns(pattern_contents: List[str]) -> List[List[float]]:
        """Fallback: return simple token-based vectors."""
        vectors = []
        for content in pattern_contents:
            # Simple keyword-based vector (bag of words)
            words = content.lower().split()
            # Create a simple hash-based vector
            vector = [hash(word) % 100 / 100.0 for word in words[:50]]  # Limit to 50 words
            # Pad or truncate to fixed size
            while len(vector) < 100:
                vector.append(0.0)
            vectors.append(vector[:100])
        return vectors
    
    @staticmethod
    def similarity_search(query: str, pattern_vectors: List[List[float]], top_k: int = 5) -> List[int]:
        """Fallback: simple keyword overlap similarity."""
        query_words = set(query.lower().split())
        
        # Simple overlap-based scoring
        scores = []
        for i, vector in enumerate(pattern_vectors):
            # This is a very basic fallback - in real implementation we'd need pattern content
            score = len(query_words) * 0.1  # Basic fallback score
            scores.append((score, i))
        
        scores.sort(reverse=True)
        return [idx for _, idx in scores[:top_k]]


class CachetoolsFallback:
    """Fallback implementation for caching without cachetools."""
    
    def __init__(self, maxsize: int = 128, ttl: int = 300):
        self.maxsize = maxsize
        self.ttl = ttl
        self._cache: Dict[str, Dict[str, Any]] = {}
    
    def get(self, key: str, default: Any = None) -> Any:
        """Get item from cache with TTL check."""
        import time
        
        if key not in self._cache:
            return default
        
        entry = self._cache[key]
        if time.time() - entry["timestamp"] > self.ttl:
            del self._cache[key]
            return default
        
        return entry["value"]
    
    def __setitem__(self, key: str, value: Any) -> None:
        """Set item in cache with timestamp."""
        import time
        
        # Simple LRU: remove oldest if at capacity
        if len(self._cache) >= self.maxsize:
            oldest_key = min(self._cache.keys(), key=lambda k: self._cache[k]["timestamp"])
            del self._cache[oldest_key]
        
        self._cache[key] = {
            "value": value,
            "timestamp": time.time()
        }
    
    def __getitem__(self, key: str) -> Any:
        """Get item from cache (raises KeyError if not found or expired)."""
        value = self.get(key)
        if value is None and key not in self._cache:
            raise KeyError(key)
        return value
    
    def __contains__(self, key: str) -> bool:
        """Check if key exists and is not expired."""
        return self.get(key) is not None


def get_cache_implementation(maxsize: int = 128, ttl: int = 300):
    """Get best available cache implementation."""
    try:
        from cachetools import TTLCache
        return TTLCache(maxsize=maxsize, ttl=ttl)
    except ImportError:
        return CachetoolsFallback(maxsize=maxsize, ttl=ttl)


def safe_import_with_fallback(module_name: str, fallback_class: type = None):
    """
    Import module with fallback class if import fails.
    
    Args:
        module_name: Name of module to import
        fallback_class: Fallback class to use if import fails
        
    Returns:
        Imported module or fallback class instance
    """
    from services.dependencies import check_and_install_if_missing
    
    if check_and_install_if_missing(module_name, module_name):
        try:
            import importlib
            return importlib.import_module(module_name)
        except ImportError:
            pass
    
    if fallback_class:
        return fallback_class()
    
    return None
