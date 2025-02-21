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

  const [fundId, setFundId] = React.useState<string>("")
  const [contributions, setContributions] = React.useState<Contribution[]>([])

  React.useEffect(() => {
    async function fetchFundId() {
      const fundId = (await params).fundId
      setFundId(fundId)
    }
    fetchFundId()
  }, [])

  React.useEffect(() => {
    async function fetchContributions() {
      const token = await getAccessToken();
      const res = await fetch(`/api/contributions?fundId=${fundId}`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
      })
      const data = await res.json()
      setContributions(data)
    }
    if (fundId) {
      fetchContributions()
    }
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
