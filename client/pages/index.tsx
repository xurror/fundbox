import Head from 'next/head'
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';

import Steps from '../components/Steps'

export default function Home() {
  return (
    <div className="h-screen min-h-screen flex flex-col justify-center items-center px-[200px] py-16 bg-primary-light">
      <Head>
        <title>Fundbox app</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col flex-1 bg-basic-200 w-full rounded-3xl shadow-md shadow px-20">
        <div className='py-14 border-b border-basic-500'>
          <h1 className="text-4xl font-bold text-secondary-dark">Create a Fund</h1>
          <p className="text-md text-basic-600 pt-1">Create a fund, copy and share the link with your friends!</p>
        </div>

        <div className='flex flex-1'>
          <div className='w-1/3 border-r border-basic-500 px-5 py-10'>
            <Steps />
          </div>
          <div className='w-2/3 p-10'>
            inputs
          </div>
        </div>
      </main>
    </div>
  )
}
