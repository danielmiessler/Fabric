#!/usr/bin/env python3
"""
Standalone dependency installer for Fabric Pattern Studio.
Run this script to install all required and optional dependencies.
"""
import sys
import os
from pathlib import Path

# Add project to path
script_dir = Path(__file__).parent
sys.path.insert(0, str(script_dir))

def main():
    """Main installation function."""
    print("🎭 Fabric Pattern Studio - Dependency Installer")
    print("=" * 50)
    
    try:
        from services.dependencies import get_dependency_manager
        
        manager = get_dependency_manager()
        
        # Check current dependency status
        print("\n📋 Checking current dependencies...")
        report = manager.check_all_dependencies()
        
        # Report status
        print(f"\n✅ Core dependencies: {len(report['core_status'])} checked")
        print(f"✨ Optional dependencies: {len(report['optional_status'])} checked")
        
        missing_core = report['missing_core']
        missing_optional = report['missing_optional']
        
        if not missing_core and not missing_optional:
            print("\n🎉 All dependencies are already installed!")
            return True
        
        # Show what needs to be installed
        if missing_core:
            print(f"\n❌ Missing CORE dependencies ({len(missing_core)}):")
            for spec in missing_core:
                print(f"  • {spec.package}")
        
        if missing_optional:
            print(f"\n⚠️ Missing OPTIONAL dependencies ({len(missing_optional)}):")
            for spec in missing_optional:
                print(f"  • {spec.package} - {spec.fallback_message}")
        
        # Ask for user confirmation
        if len(sys.argv) > 1 and "--auto" in sys.argv:
            install_all = True
            print("\n🤖 Auto-install mode enabled")
        else:
            response = input("\n❓ Install missing dependencies? [y/N]: ").strip().lower()
            install_all = response in ['y', 'yes']
        
        if not install_all:
            print("❌ Installation cancelled by user")
            
            # Show manual installation commands
            print("\n📝 Manual installation commands:")
            print("pip install -r requirements.txt")
            print("\n Or install individually:")
            for cmd in report['installation_commands']:
                print(f"  {cmd}")
            
            return False
        
        # Install missing dependencies
        all_missing = missing_core + missing_optional
        
        print(f"\n📦 Installing {len(all_missing)} dependencies...")
        results = manager.install_missing_dependencies(all_missing)
        
        # Report results
        successful = sum(1 for success in results.values() if success)
        failed = len(results) - successful
        
        print(f"\n📊 Installation Results:")
        print(f"  ✅ Successful: {successful}")
        print(f"  ❌ Failed: {failed}")
        
        if failed > 0:
            print(f"\n❌ Some installations failed:")
            for package, success in results.items():
                if not success:
                    print(f"  • {package}")
            print(f"\n💡 Try installing manually: pip install -r requirements.txt")
            return False
        
        print(f"\n🎉 All dependencies installed successfully!")
        print(f"You can now run: ./run.sh")
        return True
        
    except Exception as e:
        print(f"\n💥 Installation failed with error: {e}")
        print(f"💡 Try manual installation: pip install -r requirements.txt")
        return False


if __name__ == "__main__":
    try:
        success = main()
        sys.exit(0 if success else 1)
    except KeyboardInterrupt:
        print("\n\n⚠️ Installation interrupted by user")
        sys.exit(1)
    except Exception as e:
        print(f"\n💥 Unexpected error: {e}")
        sys.exit(1)
