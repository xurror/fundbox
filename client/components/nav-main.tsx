"use client"

import { ChevronRight, type LucideIcon } from "lucide-react"

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"
import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar"
import React from "react"
import { getAccessToken } from "@auth0/nextjs-auth0"
import Link from "next/link"

type Item = {
  title: string
  path: string
  icon?: LucideIcon
  isActive?: boolean
}

export function NavMain({
  items,
}: {
  items: Item[]
}) {
  return (
    <SidebarGroup>
      <SidebarGroupLabel>Platform</SidebarGroupLabel>
      <SidebarMenu>
        {items.map((item) => (
          <MenuContent key={item.title} item={item} />
        ))}
      </SidebarMenu>
    </SidebarGroup>
  )
}

function MenuContent({ item }: { item: Item }) {
  const [funds, setFunds] = React.useState<{
    id: string
    name: string
    targetAmount: number
  }[]>([])

  React.useEffect(() => {
    if (!item.path) {
      return
    }

    (async () => {
      const res = await fetch(`/api${item.path}`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${await getAccessToken()}`
        },
      })
      const data = await res.json()
      setFunds(data)
    })();
  }, [item.path])

  return (
    <Collapsible
      asChild
      defaultOpen={item.isActive}
      className="group/collapsible"
    >
      <SidebarMenuItem>
        <CollapsibleTrigger asChild>
          <SidebarMenuButton tooltip={item.title}>
            {item.icon && <item.icon />}
            <span>{item.title}</span>
            <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
          </SidebarMenuButton>
        </CollapsibleTrigger>
        <CollapsibleContent>
          <SidebarMenuSub>
            {funds?.map((subItem) => (
              <SidebarMenuSubItem key={subItem.name}>
                <SidebarMenuSubButton asChild>
                  <Link href={`/dashboard/funds/${subItem.id}`}>
                    <span>{subItem.name}</span>
                  </Link>
                </SidebarMenuSubButton>
              </SidebarMenuSubItem>
            ))}
          </SidebarMenuSub>
        </CollapsibleContent>
      </SidebarMenuItem>
    </Collapsible>
  )
}
