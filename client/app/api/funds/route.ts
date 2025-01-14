// pages/api/products.js
import { getAccessToken, getSession, withApiAuthRequired } from "@auth0/nextjs-auth0";
import { NextRequest, NextResponse } from "next/server";

const POST = withApiAuthRequired(async function POST(req: NextRequest) {
  const res = new NextResponse();
  const session = await getSession();
  console.log({ session });
  const response = await fetch("http://localhost:8080/api/hello", {
    headers: {
      Authorization: `${session?.token_type} ${session?.idToken}`,
    },
  });
  const products = await response.json();

  return NextResponse.json(products, res);
  // return NextResponse.json({ foo: 'bar' }, res);
});

export { POST };

// export default withApiAuthRequired(
// export async function POST(request: NextRequest) {
//   // If your access token is expired and you have a refresh token
//   // `getAccessToken` will fetch you a new one using the `refresh_token` grant
//   const nextResponse = new NextResponse()

//   const { accessToken } = await getAccessToken(request, nextResponse, {
//     scopes: [],
//   });

//   console.log({ accessToken });
//   const response = await fetch("http://localhost:8080/api/hello", {
//     headers: {
//       Authorization: `Bearer ${accessToken}`,
//     },
//   });
//   const products = await response.json();

//   return NextResponse.json(products, { status: 200 });
//   // nextResponse.status(200).json(products);
// }
// );
