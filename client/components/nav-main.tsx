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
import Link from "next/link"
import { useQuery } from "@tanstack/react-query"

type Item = {
  title: string
  path: string
  icon?: LucideIcon
  isActive?: boolean
}

type Fund = {
  id: string
  name: string
  targetAmount: number
  createdAt: string
  updatedAt: string
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

function useFunds(path: string) {
  return useQuery({
    queryKey: ['funds'],
    queryFn: async (): Promise<Array<Fund>> => {
      const response = await fetch(`/api${path}`, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
      return await response.json()
    },
  })
}

function MenuContent({ item }: { item: Item }) {
  const { data } = useFunds(item.path)

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
            {data && data.length > 0 && data?.map((subItem) => (
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
