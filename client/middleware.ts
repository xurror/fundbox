import { NextRequest, NextResponse } from "next/server"

import { auth0 } from "./lib/auth0"

// 1. Specify protected and public routes
const protectedRoutes = ['/dashboard', '/api']
const publicRoutes = ['/login', '/signup', '/']
 
export async function middleware(request: NextRequest) {

  const path = request.nextUrl.pathname
  const isProtectedRoute = protectedRoutes.includes(path)
  const isPublicRoute = publicRoutes.includes(path)
  const isAuthRoute = path.startsWith('/auth')
  const isApiRoute = path.startsWith('/api')

  const authRes = await auth0.middleware(request)

  if (isAuthRoute || isPublicRoute) {
    // skip auth for auth routes and public routes
    return authRes
  }

  const session = await auth0.getSession(request)
  if (isProtectedRoute && !session) {
    // user is not authenticated, redirect to login page
    return NextResponse.redirect(new URL("/auth/login", request.nextUrl.origin))
  }

  if (isApiRoute) {
    return NextResponse.next()
  }

  // the headers from the auth middleware should always be returned
  return authRes
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico, sitemap.xml, robots.txt (metadata files)
     */
    "/((?!_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
  ],
}
