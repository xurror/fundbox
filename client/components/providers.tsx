"use client"

import * as React from "react"
import { ThemeProvider as NextThemesProvider } from "next-themes"

export function Providers({ children }: { children: React.ReactNode }) {
  const options = { clientSecret: null };
  return (
    <NextThemesProvider
      attribute="class"
      defaultTheme="light"
      // defaultTheme="system"
      enableSystem
      disableTransitionOnChange
      enableColorScheme
    >
        {children}
    </NextThemesProvider>
  )
}
