"use client"

import * as React from "react"
import {
  GalleryVerticalEnd,
  HandCoins,
  Landmark,
} from "lucide-react"

import { NavMain } from "@/components/nav-main"
import { NavUser } from "@/components/nav-user"
import { TeamSwitcher } from "@/components/team-switcher"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar"
import { useAppUser } from "@/hooks/use-app-user"

const teams = [
  {
    name: "Community Funds",
    logo: GalleryVerticalEnd,
    plan: "Enterprise",
  },
]

const getMainNavMenu = (contributorId: string) => [
  {
    title: "My Funds",
    path: "/funds",
    icon: HandCoins,
    isActive: true,
  },
  {
    title: "My Communities",
    path: "/funds",
    params: { contributorId: contributorId },
    icon: Landmark,
  },
]

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { data: user } = useAppUser()

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <TeamSwitcher teams={teams} />
      </SidebarHeader>
      <SidebarContent>
        {user && <NavMain items={getMainNavMenu(user.id)} />}
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
