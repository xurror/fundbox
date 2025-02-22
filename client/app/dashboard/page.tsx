import { Chart1 } from "@/components/chart-1";
import { ContributionsDataTable } from "@/components/contributions/data-table";
import { auth0 } from "@/lib/auth0";

export default async function Page() {

  const res = await fetch(`${process.env.BACKEND_URL}/api/contributions`, {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${(await auth0.getSession())?.tokenSet.accessToken}`
    },
  })
  const contributions = await res.json()

  return (
    <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
      <div className="grid auto-rows-min gap-4 md:grid-cols-3">
        <div className="bg-muted/50 aspect-video rounded-xl">
          <Chart1 />
        </div>
        <div className="bg-muted/50 aspect-video rounded-xl" />
        <div className="bg-muted/50 aspect-video rounded-xl" />
      </div>
      <div className="bg-muted/50 min-h-[100vh] flex-1 rounded-xl px-4 md:min-h-min">
        <div className="container mx-auto">
          <ContributionsDataTable data={contributions} />
        </div>
      </div>
    </div>
  )
}
