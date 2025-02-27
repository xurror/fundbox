"use client"

import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { useContributions } from "@/hooks/use-contributions"
import { formatAmount } from "@/lib/formatter"

export function RecentContributions() {
  const { data: contributions } = useContributions()
  return (
    <Card className="col-span-3">
      <CardHeader>
        <CardTitle>Recent Sales</CardTitle>
        <CardDescription>
          You made {(contributions as [])?.length} sales this month.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-8">
          {contributions && contributions.slice(0, 5).map((contribution) => (
            <div className="flex items-center" key={contribution.id}>
              <Avatar className="h-9 w-9">
                <AvatarImage src={undefined} alt="Avatar" />
                <AvatarFallback>{contribution.contributorName.slice(0, 2)}</AvatarFallback>
              </Avatar>
              <div className="ml-4 space-y-1">
                <p className="text-sm font-medium leading-none">{contribution.contributorName}</p>
                <p className="text-sm text-muted-foreground">{contribution.fundName}</p>
              </div>
              <div className="ml-auto font-medium">+{formatAmount(contribution.amount)}</div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}
