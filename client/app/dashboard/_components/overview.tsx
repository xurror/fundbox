"use client"

import { Contribution } from "@/components/contributions/columns";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from "recharts"

const transformData = (rawData: Contribution[]) => {
  const months = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
  ];

  const transformed = months.map((month) => ({
    name: month,
    total: 0,
  }));

  rawData.forEach((entry) => {
    const date = new Date(entry.createdAt);
    const monthIndex = date.getMonth();
    transformed[monthIndex].total += entry.amount;
  });

  return transformed;
};

export function Overview({ data: rawData }: { data: Contribution[] }) {
  const data = transformData(rawData);
  
  return (
    <ResponsiveContainer width="100%" height={350}>
      <BarChart data={data}>
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
  )
}
