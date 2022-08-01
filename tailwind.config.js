/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{md,njk}',
  ],
  theme: {
    extend: {},
    fontFamily: {
      'serif': ['Roboto Slab', 'Arial', 'serif'],
      'sans': ['Roboto', 'Arial', 'sans-serif'],
      'mono': ['Roboto Mono', 'Arial', 'mono'],
    }
  },
  plugins: [],
}
