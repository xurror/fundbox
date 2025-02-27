"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useContributions } from "@/hooks/use-contributions";
import { Contribution } from "@/types/contribution";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from "recharts"

const transformData = (rawData: Array<Contribution>) => {
  const months = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
  ];

  const transformed = months.map((month) => ({
    name: month,
    total: 0,
  }));

  (rawData ?? []).forEach((entry) => {
    const date = new Date(entry.createdAt);
    const monthIndex = date.getMonth();
    transformed[monthIndex].total += entry.amount;
  });

  return transformed;
};

export function ContributionsOverview() {
  const { data } = useContributions()
  return (
    <Card className="col-span-4">
      <CardHeader>
        <CardTitle>Overview</CardTitle>
      </CardHeader>
      <CardContent className="pl-2">
        {data &&
          <ResponsiveContainer width="100%" height={350}>
            <BarChart data={transformData(data)}>
              <XAxis
                dataKey="name"
                stroke="#888888"
                fontSize={12}
                tickLine={false}
                axisLine={false}
              />
              <YAxis
                stroke="#888888"
                fontSize={12}
                tickLine={false}
                axisLine={false}
                tickFormatter={(value) => `$${value}`}
              />
              <Bar
                dataKey="total"
                fill="currentColor"
                radius={[4, 4, 0, 0]}
                className="fill-primary"
              />
            </BarChart>
          </ResponsiveContainer>
        }
      </CardContent>
    </Card>

  )
}
