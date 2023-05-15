/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{md,njk}',
  ],
  theme: {
    extend: {
      colors: {
        'dk-color': '#c41e3a',
        'druid-color': '#ff7c0a',
        'hunter-color': '#aad372',
        'mage-color': '#3fc7eb',
        'paladin-color': '#f48cba',
        'priest-color': '#ffffff',
        'rogue-color': '#fff468',
        'shaman-color': '#0070dd',
        'warlock-color': '#8788ee',
        'warrior-color': '#c69b6d',
      }
    },
    fontFamily: {
      'serif': ['Roboto Slab', 'Arial', 'serif'],
      'sans': ['Roboto', 'Arial', 'sans-serif'],
      'mono': ['Roboto Mono', 'Arial', 'mono'],
    }
  },
  safelist: [
    'text-dk-color',
    'text-druid-color',
    'text-hunter-color',
    'text-mage-color',
    'text-paladin-color',
    'text-priest-color',
    'text-rogue-color',
    'text-shaman-color',
    'text-warlock-color',
    'text-warrior-color',
    'bg-dk-color',
    'bg-druid-color',
    'bg-hunter-color',
    'bg-mage-color',
    'bg-paladin-color',
    'bg-priest-color',
    'bg-rogue-color',
    'bg-shaman-color',
    'bg-warlock-color',
    'bg-warrior-color',
  ],
  plugins: [],
}
