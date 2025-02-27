"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useContributions } from "@/hooks/use-contributions";
import { formatAmount } from "@/lib/formatter";
import React from "react";

export function TotalContributions() {
  const { data: contributions } = useContributions()
  const [total, setTotal] = React.useState(0)
  const [percentage, setPercentage] = React.useState(0)

  React.useEffect(() => {
    if (contributions) {
      const total = contributions.reduce((acc, curr) => acc + curr.amount, 0)
      setTotal(total)
    }
  }, [contributions])

  React.useEffect(() => {
    if (contributions) {
      const total = contributions.reduce((acc, curr) => acc + curr.amount, 0)
      const lastMonthTotal = contributions
        .filter((contribution) => {
          const date = new Date(contribution.createdAt)
          const lastMonth = new Date()
          lastMonth.setMonth(lastMonth.getMonth() - 1)
          return date.getMonth() === lastMonth.getMonth()
        })
        .reduce((acc, curr) => acc + curr.amount, 0)
      
      const percentage = lastMonthTotal == 0 ? 100 : ((total - lastMonthTotal) / lastMonthTotal) * 100
      setPercentage(percentage)
    }
  }, [contributions])

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">
          Total Contributions
        </CardTitle>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          className="h-4 w-4 text-muted-foreground"
        >
          <path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6" />
        </svg>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">{formatAmount(total)}</div>
        <p className="text-xs text-muted-foreground">
          +{percentage}% from last month
        </p>
      </CardContent>
    </Card>
  )
}
