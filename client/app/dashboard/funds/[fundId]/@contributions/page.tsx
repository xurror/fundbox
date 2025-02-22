import { ContributionsDataTable } from "@/components/contributions/data-table";
import { auth0 } from "@/lib/auth0";

export default async function Page({ params }: { params: Promise<{ fundId: string }> }) {
  const fundId = (await params).fundId
  console.log(fundId)
  const res = await fetch(`http://localhost:8080/api/contributions?fundId=${fundId}`, {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${(await auth0.getSession())?.tokenSet.accessToken}`
    },
  })
  const contributions = await res.json()

  return (
    <div className="container mx-auto">
      <ContributionsDataTable data={contributions} />
    </div>
  )
}
