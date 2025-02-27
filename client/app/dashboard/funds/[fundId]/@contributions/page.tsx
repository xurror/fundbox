"use client"

import { ContributionsDataTable } from "@/components/contributions/data-table";
import React from "react";

export default function Page({ params }: { params: Promise<{ fundId: string }> }) {
  const fundId = React.use(params).fundId

  return <ContributionsDataTable fundId={fundId} />
}
