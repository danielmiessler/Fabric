"""
Provider service for managing AI vendors and models.
Implements the vendor/model architecture from Fabric.
"""
import os
import subprocess
from dataclasses import dataclass, field
from typing import Dict, List, Optional, Tuple
from pathlib import Path
import json

from utils.logging import logger
from utils.errors import ValidationError


@dataclass
class VendorConfig:
    """Vendor configuration with required settings."""
    name: str
    api_key: Optional[str] = None
    base_url: Optional[str] = None
    enabled: bool = False
    models: List[str] = field(default_factory=list)
    env_var_name: Optional[str] = None  # e.g., OPENAI_API_KEY


@dataclass
class ModelSpec:
    """Model specification with vendor context."""
    vendor: str
    model: str
    display_name: str
    context_length: int = 2048
    
    @property
    def full_name(self) -> str:
        """Get full model name in vendor/model format."""
        return f"{self.vendor}/{self.model}"


@dataclass
class ConfigQuestion:
    """Configuration question for vendor setup."""
    key: str
    prompt: str
    required: bool = True
    is_secret: bool = False
    default: Optional[str] = None


class ProviderService:
    """Encapsulates vendor operations and model management."""
    
    # Known vendor configurations
    VENDOR_CONFIGS = {
        "openai": {
            "env_var": "OPENAI_API_KEY",
            "base_url_var": "OPENAI_BASE_URL",
            "models_prefix": ["gpt-"],
        },
        "anthropic": {
            "env_var": "ANTHROPIC_API_KEY", 
            "base_url_var": "ANTHROPIC_BASE_URL",
            "models_prefix": ["claude-"],
        },
        "ollama": {
            "env_var": None,  # No API key needed
            "base_url_var": "OLLAMA_URL",
            "models_prefix": ["llama", "mistral", "mixtral", "codellama"],
        },
        "azure": {
            "env_var": "AZURE_OPENAI_API_KEY",
            "base_url_var": "AZURE_OPENAI_ENDPOINT",
            "models_prefix": ["gpt-", "azure-"],
        },
        "gemini": {
            "env_var": "GEMINI_API_KEY",
            "base_url_var": "GEMINI_BASE_URL",
            "models_prefix": ["gemini-"],
        },
        "perplexity": {
            "env_var": "PERPLEXITY_API_KEY",
            "base_url_var": "PERPLEXITY_BASE_URL",
            "models_prefix": ["pplx-"],
        },
        "groq": {
            "env_var": "GROQ_API_KEY",
            "base_url_var": "GROQ_BASE_URL",
            "models_prefix": ["llama", "mixtral"],
        },
        "openrouter": {
            "env_var": "OPENROUTER_API_KEY",
            "base_url_var": "OPENROUTER_API_BASE_URL",
            "models_prefix": ["openrouter/", "or/"],
        },
    }
    
    @classmethod
    def load_vendor_configs(cls) -> Dict[str, VendorConfig]:
        """Load vendor configs from .env and validate."""
        configs = {}
        
        # Read from Fabric's .env file (use the actual user's home)
        # Try the actual user's home directory first
        import os
        actual_home = os.path.expanduser("~")
        env_path = Path(actual_home) / ".config" / "fabric" / ".env"
        env_vars = cls._read_env_file(env_path)
        
        # Also check system environment
        for key, value in os.environ.items():
            if key not in env_vars:
                env_vars[key] = value
        
        # Process each known vendor
        for vendor_name, vendor_info in cls.VENDOR_CONFIGS.items():
            config = VendorConfig(name=vendor_name)
            
            # Check for API key
            if vendor_info["env_var"]:
                api_key = env_vars.get(vendor_info["env_var"])
                if api_key:
                    config.api_key = api_key
                    config.enabled = True
                    config.env_var_name = vendor_info["env_var"]
            elif vendor_name == "ollama":
                # Skip Ollama if not available
                config.enabled = False
            
            # Check for base URL
            if vendor_info.get("base_url_var"):
                base_url = env_vars.get(vendor_info["base_url_var"])
                if base_url:
                    config.base_url = base_url
            
            # Get models for this vendor
            config.models = cls._get_vendor_models(vendor_name, env_vars)
            
            configs[vendor_name] = config
        
        return configs
    
    @classmethod
    def list_available_models(cls, vendor: str = None) -> List[ModelSpec]:
        """Query vendor for supported models with proper fabric CLI parsing."""
        models = []
        
        try:
            # Use fabric CLI to get models
            cmd = ["fabric", "--listmodels"]
            result = subprocess.run(cmd, capture_output=True, text=True, timeout=30)
            
            if result.returncode == 0 and result.stdout:
                models = cls._parse_fabric_models_output(result.stdout, vendor)
            else:
                logger.warning(f"fabric --listmodels failed: {result.stderr}")
                
        except subprocess.TimeoutExpired:
            logger.warning("Timeout while fetching models")
        except Exception as e:
            logger.error(f"Error fetching models: {e}")
        
        return models
    
    @classmethod
    def _parse_fabric_models_output(cls, output: str, filter_vendor: str = None) -> List[ModelSpec]:
        """Parse the structured output from fabric --listmodels."""
        import re
        models = []
        lines = output.strip().split('\n')
        current_vendor = None
        
        for line in lines:
            line = line.strip()
            if not line:
                continue
                
            # Skip header and error messages
            if ('Available models:' in line or 
                'Ollama Get' in line or 
                'dial tcp' in line or
                'connect: connection refused' in line):
                continue
                
            # Check if this is a vendor header (not indented, not numbered)
            if (not line.startswith('\t') and 
                not line.startswith('[') and 
                not line.startswith('  [')):
                # This is a vendor name like "OpenAI", "Gemini", etc.
                current_vendor = line.lower()
                continue
                
            # Parse numbered model entries: [1] model-name or \t[1]\tmodel-name
            if current_vendor and ('[' in line):
                # Handle both formats: "[1] model" and "\t[1]\tmodel"
                match = re.search(r'\[(\d+)\]\s*(.+)', line)
                if match:
                    model_name = match.group(2).strip()
                    
                    # Filter by vendor if specified
                    if filter_vendor and current_vendor != filter_vendor.lower():
                        continue
                        
                    # Create model spec
                    spec = ModelSpec(
                        vendor=current_vendor,
                        model=model_name,
                        display_name=model_name,
                        context_length=cls._get_model_context_length(model_name)
                    )
                    models.append(spec)
        
        logger.info(f"Parsed {len(models)} models from fabric CLI output")
        return models
    
    @classmethod
    def validate_model_access(cls, vendor: str, model: str) -> bool:
        """Test if model is accessible with current config."""
        configs = cls.load_vendor_configs()
        
        # Check if vendor is configured
        if vendor not in configs:
            return False
        
        vendor_config = configs[vendor]
        if not vendor_config.enabled:
            return False
        
        # For API-based vendors, check if key exists
        if vendor != "ollama" and not vendor_config.api_key:
            return False
        
        # Check if model is in vendor's model list
        available_models = cls.list_available_models(vendor)
        model_names = [m.model for m in available_models]
        
        return model in model_names
    
    @classmethod
    def get_vendor_questions(cls, vendor: str) -> List[ConfigQuestion]:
        """Get configuration questions for vendor setup."""
        questions = []
        
        vendor_info = cls.VENDOR_CONFIGS.get(vendor)
        if not vendor_info:
            return questions
        
        # API key question
        if vendor_info.get("env_var"):
            questions.append(ConfigQuestion(
                key="api_key",
                prompt=f"Enter your {vendor.upper()} API key",
                required=True,
                is_secret=True
            ))
        
        # Base URL question
        if vendor_info.get("base_url_var"):
            default_url = None
            if vendor == "ollama":
                default_url = "http://localhost:11434"
            
            questions.append(ConfigQuestion(
                key="base_url",
                prompt=f"Enter {vendor.upper()} base URL",
                required=vendor == "ollama",
                is_secret=False,
                default=default_url
            ))
        
        # Azure specific questions
        if vendor == "azure":
            questions.append(ConfigQuestion(
                key="deployment_name",
                prompt="Enter your Azure OpenAI deployment name",
                required=True,
                is_secret=False
            ))
        
        return questions
    
    @classmethod
    def detect_vendor_from_model(cls, model_name: str) -> str:
        """Detect vendor from model name."""
        return cls._detect_vendor_from_model(model_name)
    
    @staticmethod
    def _read_env_file(path: Path) -> Dict[str, str]:
        """Read environment variables from .env file."""
        env_vars = {}
        
        if not path.exists():
            return env_vars
        
        try:
            with open(path, 'r') as f:
                for line in f:
                    line = line.strip()
                    if line and not line.startswith('#') and '=' in line:
                        key, value = line.split('=', 1)
                        # Remove quotes if present
                        value = value.strip().strip('"').strip("'")
                        env_vars[key] = value
        except Exception as e:
            logger.warning(f"Error reading .env file: {e}")
        
        return env_vars
    
    @classmethod
    def _get_vendor_models(cls, vendor: str, env_vars: Dict[str, str]) -> List[str]:
        """Get models for a specific vendor from environment."""
        models = []
        
        # Check for vendor-specific model list in env
        model_list_key = f"{vendor.upper()}_MODELS"
        if model_list_key in env_vars:
            models = [m.strip() for m in env_vars[model_list_key].split(',')]
        
        return models
    
    @classmethod
    def _detect_vendor_from_model(cls, model_name: str) -> str:
        """Detect vendor from model name patterns."""
        model_lower = model_name.lower()
        
        # Direct vendor name mappings (case-insensitive)
        vendor_mappings = {
            "openai": "openai",
            "anthropic": "anthropic", 
            "ollama": "ollama",
            "azure": "azure",
            "gemini": "gemini",
            "perplexity": "perplexity",
            "groq": "groq",
            "openrouter": "openrouter"
        }
        
        # Check for exact vendor name match first (for DEFAULT_VENDOR)
        for vendor_name, vendor_key in vendor_mappings.items():
            if model_lower == vendor_name:
                return vendor_key
        
        # Then check model name prefixes
        for vendor, info in cls.VENDOR_CONFIGS.items():
            for prefix in info.get("models_prefix", []):
                if model_lower.startswith(prefix):
                    return vendor
        
        return "unknown"
    
    @staticmethod
    def _get_model_context_length(model_name: str) -> int:
        """Get context length for known models."""
        # Known context lengths
        context_lengths = {
            "gpt-4": 8192,
            "gpt-4-turbo": 128000,
            "gpt-4o": 128000,
            "gpt-3.5-turbo": 4096,
            "claude-3-opus": 200000,
            "claude-3-sonnet": 200000,
            "claude-2.1": 200000,
            "gemini-pro": 32000,
            "mixtral": 32000,
        }
        
        # Check exact match first
        if model_name in context_lengths:
            return context_lengths[model_name]
        
        # Check partial matches
        for key, length in context_lengths.items():
            if key in model_name.lower():
                return length
        
        # Default context length
        return 2048