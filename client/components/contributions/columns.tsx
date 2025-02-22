"use client"

import { DataTableColumnHeader } from "./data-table-column-header"
import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox"
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuLabel
} from "@/components/ui/dropdown-menu"
import { ColumnDef } from "@tanstack/react-table"
import { MoreHorizontal } from "lucide-react"

export type Contribution = {
  id: string
  fundName: string
  contributorName: string
  amount: number
  createdAt: string
}

export const columns: ColumnDef<Contribution>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "fundName",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Fund Name" />,
  },
  {
    accessorKey: "contributorName",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Contributor Name" />,
  },
  {
    accessorKey: "amount",
    header: ({ column }) => <DataTableColumnHeader column={column} title="Amount" className="justify-end" />,
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("amount"))
      const formatted = new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "USD",
      }).format(amount)

      return <div className="text-right font-medium">{formatted}</div>
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const payment = row.original
      return (
        <div className="size-2 shrink-0">
         <DropdownMenu>
           <DropdownMenuTrigger asChild>
             <Button variant="ghost" className="size-4">
               <span className="sr-only">Open menu</span>
               <MoreHorizontal className="size-4 shrink-0" />
             </Button>
           </DropdownMenuTrigger>
           <DropdownMenuContent align="end">
             <DropdownMenuLabel>Actions</DropdownMenuLabel>
             <DropdownMenuItem
               onClick={() => navigator.clipboard.writeText(payment.id)}
             >
               Copy payment ID
             </DropdownMenuItem>
             <DropdownMenuSeparator />
             <DropdownMenuItem>View customer</DropdownMenuItem>
             <DropdownMenuItem>View payment details</DropdownMenuItem>
           </DropdownMenuContent>
         </DropdownMenu>
        </div>
      )
    },
  },
]
