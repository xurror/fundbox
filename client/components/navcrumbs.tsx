"use client"

import {
  Breadcrumb,
  BreadcrumbList,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbSeparator,
  BreadcrumbPage
} from "@/components/ui/breadcrumb";
import { usePathname } from "next/navigation";
import React, { JSX } from "react";

async function fetchFund(fundId: string) {
  const res = await fetch(`/api/funds/${fundId}`, {
    headers: { 'Content-Type': 'application/json' },
  })
  return res.json()
}

export default function Navcrumbs() {
  const pathname = usePathname()

  const [breadcrumbs, setBreadcrumbs] = React.useState<JSX.Element[]>([]);

  React.useEffect(() => {
    (async () => {
      const paths = pathname.split("/").filter(Boolean)
      paths.shift() // drop the "dashboard" part

      // group paths in pairs of 2, if the end is odd, drop the last one
      const pairs = paths.reduce((acc: string[][], path, i) => {
        if (i % 2 === 0) {
          acc.push([path])
        } else {
          acc[acc.length - 1].push(path)
        }
        return acc
      }, [])

      const items = await Promise.all(
        pairs.map(async (pair, index) => {
          const fund = await fetchFund(pair[1]);
          const isLast = index === pairs.length - 1;
          return (
            <React.Fragment key={pair[1]}>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                {isLast ? (
                  <BreadcrumbPage>{fund.name}</BreadcrumbPage>
                ) : (
                  <BreadcrumbLink href={`/${pair.join("/")}`}>{fund.name}</BreadcrumbLink>
                )}
              </BreadcrumbItem>
              {!isLast && <BreadcrumbSeparator />}
            </React.Fragment>
          );
        })
      );
      setBreadcrumbs(items);
    })();
  }, [pathname]);

  return (
    <Breadcrumb>
      <BreadcrumbList>
        <BreadcrumbItem className="hidden md:block">
          <BreadcrumbLink href="/dashboard">
            Dashboard
          </BreadcrumbLink>
        </BreadcrumbItem>
        {breadcrumbs}
      </BreadcrumbList>
    </Breadcrumb>
  );
}
