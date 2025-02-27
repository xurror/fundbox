import { Metadata } from "next"

import { RecentContributions } from "./_components/recent-contributions"
import { NewFundForm } from "@/components/new-fund-form"
import { NewContributionForm } from "@/components/new-contribution-form"
import { ContributionsOverview } from "@/components/charts/contributions-overview"
import { TotalContributions } from "@/components/charts/total-contributions"
import { ContributionsDataTable } from "@/components/contributions/data-table"

export const metadata: Metadata = {
  title: "Dashboard",
}

export default async function Page() {
  
  return (
    <div className="flex-col md:flex">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold tracking-tight -mb-1">Dashboard</h2>
          <div className="flex items-center space-x-2">
            <NewFundForm />
            <NewContributionForm />
          </div>
        </div>
        <div className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <TotalContributions />
            
            <div className="bg-muted/50 aspect-video rounded-xl" />
            <div className="bg-muted/50 aspect-video rounded-xl" />
            <div className="bg-muted/50 aspect-video rounded-xl" />
          </div>
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-7">
            <ContributionsOverview />
            <RecentContributions />
          </div>
          <ContributionsDataTable />
        </div>
      </div>
    </div>
  )
}
