import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import Countries from "../components/Countries";

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Fundbox app</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <a href="https://nextjs.org">FundBox!</a>
        </h1>

        <Countries />
      </main>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by Humans 😉
        </a>
      </footer>
    </div>
  )
}
