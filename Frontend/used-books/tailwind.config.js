/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}"
  ],
  theme: {
    extend: {
      colors: {
				primary: '#31AAE0',
				secondary: '#FA8A00',
				'gray-custom': '#14445D',
				'green-custom': '#25B99E',
				'red-custom': '#FA8072',
				danger: '#EA4735',
			},
      backgroundImage:{
          'book': "url('/public/buku.png')"
      },
    },
  },
  plugins: [],
}
