/** @type {import('tailwindcss').Config} */
export default {
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: "#191F4D",
          secondary: "#ffffff",
          error: "#ef4444",
          "base-100": "#ffffff",
          "--rounded-btn": "0.25rem",
          "--rounded-btn-secondary": "0.25rem",
          "--btn-text-case": "titlecase",
        },
      },
    ],
  },
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    fontFamily: {
      sans: [
        "IBM Plex Sans, sans-serif",
        {
          // fontFeatureSettings: '"cv11", "ss01"',
          // fontVariationSettings: '"opsz" 32'
        },
      ],
    },
    extend: {
      colors: {
        brand: {
          200: "#191f4d",
        },
        blue: {
          150: "#2196f3",
          250: "#0E4DDD",
        },
        slate: {
          0: "#ffffff",
          50: "#f8fafc",
          100: "#f1f5f9",
          150: "#eaeef3",
          200: "#e2e8f0",
          300: "#cbd5e1",
          400: "#94a3b8",
          500: "#64748b",
          600: "#475569",
          700: "#334155",
          800: "#1e293b",
          900: "#0f172a",
          950: "#020617",
        },
        orange: {
          50: "#fff3e0",
          100: "#fce4c0",
          200: "#ffce84",
        },
      },
    },
    fontSize: {
      sm: "12px",
      base: "14px",
      lg: "16px",
      xl: "18px",
      "2xl": "20px",
      "3xl": "24px",
      "4xl": "32px",
      "5xl": "48px",
      "6xl": "64px",
    },
    screens: {
      sm: "640px",
      md: "768px",
      lg: "1024px",
      xl: "1280px",
      "2xl": "2560px",
      "4xl": "3100px",
    },
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("daisyui"),
  ],
};
