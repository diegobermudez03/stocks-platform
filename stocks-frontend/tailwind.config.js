/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        custom: ['"Segoe UI"', 'sans-serif'],
      },
    },
  },
  plugins: [],
}

