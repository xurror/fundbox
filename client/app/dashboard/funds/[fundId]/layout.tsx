import { NewContributionForm } from "@/components/new-contribution-form"
import { NewFundForm } from "@/components/new-fund-form"

export default async function Layout({
  charts,
  contributions,
}: Readonly<{
  charts: React.ReactNode
  contributions: React.ReactNode
}>) {
  return (
    <div className="flex-col md:flex">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <div className="flex items-center justify-between space-y-2">
          <h2 className="text-3xl font-bold capitalize tracking-tight -mb-1">Dashboard</h2>
          <div className="flex items-center space-x-2">
            <NewFundForm />
            <NewContributionForm />
          </div>
        </div>
        <>
          <div className="flex flex-1 flex-col gap-4 py-4 pt-0">
            <div className="grid auto-rows-min gap-4 md:grid-cols-3">
              <div className="bg-muted/50 aspect-video rounded-xl">
                {charts}
              </div>
              <div className="bg-muted/50 aspect-video rounded-xl" />
              <div className="bg-muted/50 aspect-video rounded-xl" />
            </div>
            <>
              {contributions}
            </>
          </div>
        </>
      </div>
    </div>

  )
}
