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
      'basic-700': '#F3F2FE',
      'dark-100': '#22173F',
      'dark-200': '#1D1D1F',
      'primary': '#8B71CD',
      'primary-light': '#C6AAE090',
      'secondary': '#220DF6',
      'secondary-dark': '#412DB3',
      'success': '#25C452',
      'danger': '#0054B0',
      'warning': '#EDE323',
      'info': '#FF433D',

      'blue-100': '#4C51C6',
      'blue-10': '#4C51C61A',
      'dark_blue-100': '#12033A',
      'dark_blue-80': '#12033ACC',
      'dark_blue-60': '#04041599',
      'dark_blue-40': '#12033A66',
      'dark_blue-20': '#12033A33',
      'dark_blue-10': '#12033A1A',
      'light-100': '#F1F3FA',
      'white': '#FFFFFF',
      'white-80': '#FFFFFFCC',
      'white-10': '#FFFFFF1A',
      'grey_sub_text': '#686873',
      'grey_disabled': '#CFCFDB',
    },
  },
  plugins: [],
}

