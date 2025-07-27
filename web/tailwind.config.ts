import { join } from 'path';
import type { Config } from 'tailwindcss';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import { skeleton } from '@skeletonlabs/tw-plugin';
import { myCustomTheme } from './my-custom-theme.ts'

export default {
  darkMode: 'class',
  content: [
    './src/**/*.{html,js,svelte,svx,md,ts}',
    join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte,ts,svx,md}')
  ],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
      },
      fontFamily: {
        mono: ['Fira Code', 'monospace'],
      },
      animation: {
        fabGradient: 'fabGradient 15s ease infinite',
        blink: 'blink 1s step-end infinite',
      },
      keyframes: {
        fabGradient: {
          '0%, 100%': { 'background-size': '200% 200%, background-position: left center' },
          '50%': { 'background-size': '200% 200%, background-position: right center' },
        },
        blink: {
          '0%, 100%': { opacity: '1' },
          '50%': { opacity: '0' },
        },
      },
      typography: {
        DEFAULT: {
          css: {
            'code::before': {
              content: '""'
            },
            'code::after': {
              content: '""'
            }
          }
        }
      }
    },
  },
  plugins: [
    forms,
    typography,
    skeleton({
      themes: {
        preset: [
          {
            name: 'skeleton',
            enhancements: true
          },
          {
            name: 'modern',
            enhancements: true
          },
          {
            name: 'crimson',
            enhancements: true
          },
          {
            name: 'hamlindigo',
            enhancements: true
          },
          {
            name: 'gold-nouveau',
            enhancements: true
          },
          {
            name: 'seafoam',
            enhancements: true
          },
          {
            name: 'rocket',
            enhancements: true
          },
          {
            name: 'sahara',
            enhancements: true
          },
          {
            name: 'wintry',
            enhancements: true
          },
          {
            name: 'vintage',
            enhancements: true
          },
        ],
        custom: [
          myCustomTheme
        ]
      }
    })
  ]
}
