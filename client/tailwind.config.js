/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      'basic-100': '#ffffff',
      'basic-200': '#FAFAFA',
      'basic-300': '#F5F5F5',
      'basic-400': '#EEEEEE',
      'basic-500': '#E0E0E0',
      'basic-600': '#6F767E',
      'basic-700': '#1D1D1F',
      'primary': '#0072F0',
      'primary-tint': '#0054B0',
      'success': '#25C452',
      'danger': '#0054B0',
      'warning': '#EDE323',
      'info': '#FF433D',
    },
  },
  plugins: [],
}

