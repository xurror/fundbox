import { Contribution } from "@/components/contributions/columns"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"

export async function RecentContributions({ data }: { data: Contribution[] }) {
  const contributions = data.slice(0, 5)
  return (
    <div className="space-y-8">
      {contributions.map((contribution) => (
        <div className="flex items-center" key={contribution.id}>
          <Avatar className="h-9 w-9">
            <AvatarImage src={undefined} alt="Avatar" />
            <AvatarFallback>{contribution.contributorName.slice(0, 2)}</AvatarFallback>
          </Avatar>
          <div className="ml-4 space-y-1">
            <p className="text-sm font-medium leading-none">{contribution.contributorName}</p>
            <p className="text-sm text-muted-foreground">{contribution.fundName}</p>
          </div>
          <div className="ml-auto font-medium">+${contribution.amount}</div>
        </div>
      ))}
    </div>
  )
}
