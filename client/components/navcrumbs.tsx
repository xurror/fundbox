"use client"

import {
  Breadcrumb,
  BreadcrumbList,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbSeparator,
  BreadcrumbPage
} from "@/components/ui/breadcrumb";
import { useFund } from "@/hooks/use-funds";
import { usePathname } from "next/navigation";
import React, { JSX } from "react";

function Crumb({ fundId, path, isLast }: { fundId: string, path: string, isLast: boolean }) {
  const { data: fund } = useFund({ id: fundId });
  return (
    <React.Fragment>
      <BreadcrumbSeparator className="hidden md:block" />
      <BreadcrumbItem>
        {isLast ? (
          <BreadcrumbPage>{fund?.name}</BreadcrumbPage>
        ) : (
          <BreadcrumbLink href={path}>{fund?.name}</BreadcrumbLink>
        )}
      </BreadcrumbItem>
      {!isLast && <BreadcrumbSeparator />}
    </React.Fragment>
  );
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
          return <Crumb
            key={pair[1]}
            fundId={pair[1]}
            path={`/${pair.join("/")}`}
            isLast={index === pairs.length - 1}
          />
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
