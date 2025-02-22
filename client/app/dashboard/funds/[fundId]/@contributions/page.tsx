"use client"

import React from "react"
import { Contribution, columns } from "./columns"
import { DataTable } from "./data-table"
import { getAccessToken } from "@auth0/nextjs-auth0"

export default function Page({
  params,
}: {
  params: Promise<{ fundId: string }>
}) {

  const fundId = React.use(params).fundId
  const [contributions, setContributions] = React.useState<Contribution[]>([])

  React.useEffect(() => {
    if (!fundId) {
      return
    }

    (async () => {
      const res = await fetch(`/api/contributions?fundId=${fundId}`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${await getAccessToken()}`
        },
      })
      const data = await res.json()
      setContributions(data)
    })();
  }, [fundId])

  if (!fundId) {
    return <div>Loading...</div>
  }


  return (
    <div className="container mx-auto">
      <DataTable fundId={fundId} columns={columns} data={contributions} />
    </div>
  )
}
