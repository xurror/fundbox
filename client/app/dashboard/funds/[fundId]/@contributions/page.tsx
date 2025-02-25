"use client"

import { ContributionsDataTable } from "@/components/contributions/data-table";
import React from "react";
import { useContributions } from "@/hooks/use-contributions";
import { columns } from "./columns";

export default function Page({ params }: { params: Promise<{ fundId: string }> }) {
  const fundId = React.use(params).fundId
  const { data } = useContributions(fundId)

  return (
    <div className="container mx-auto">
      {data && <ContributionsDataTable data={data} columns={columns} />}
    </div>
  )
}
