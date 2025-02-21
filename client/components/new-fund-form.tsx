"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog"
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form"
import { Plus } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { toast } from "sonner"
import { getAccessToken } from "@auth0/nextjs-auth0"
import React from "react"

const FormSchema = z.object({
	name: z.string().min(2, {
		message: "Name must be at least 2 characters.",
	}),
	targetAmount: z.coerce.number().gt(0, {
		message: 'Target amount must be at least 1.',
	}),
})

export function NewFundForm() {
	const [open, setOpen] = React.useState(false);
	
	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: {
			name: "",
			targetAmount: 0,
		},
	})

	async function onSubmit(data: z.infer<typeof FormSchema>) {
		const token = await getAccessToken();
		const response = await fetch('/api/funds', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `Bearer ${token}`
			},
			body: JSON.stringify(data)
		})
		if (response.ok) {
			const fund = await response.json()
			toast.success("Successfully created fund", {
				description: `You've successfully created fund: ${fund.name}`,
			})
			form.reset()
			setOpen(false)
		} else {
			toast.error("Failed to create fund", {
				description: "An error occurred while creating the fund.",
			})
		}
	}


	return (
		<Dialog open={open} onOpenChange={setOpen}>
			<DialogTrigger asChild>
				<Button>
					<Plus /> Create Fund
				</Button>
			</DialogTrigger>
			<DialogContent className="sm:max-w-[512px]">
				<DialogHeader>
					<DialogTitle>Create a Fund</DialogTitle>
					<DialogDescription>
						Making an impact starts here.
					</DialogDescription>
				</DialogHeader>

				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)}>
						<div className="grid gap-4 py-4">
							<FormField
								control={form.control}
								name="name"
								render={({ field }) => (
									<FormItem className="grid grid-cols-4 items-center gap-4">
										<FormLabel className="text-right">Name</FormLabel>
										<FormControl>
											<Input placeholder="Family Retreat" className="col-span-3" {...field} />
										</FormControl>
										{/* <FormDescription className="col-span-3 col-start-2">
											This is your public display name.
										</FormDescription> */}
										<FormMessage className="col-span-3 col-start-2" />
									</FormItem>
								)}
							/>
							<FormField
								control={form.control}
								name="targetAmount"
								render={({ field }) => (
									<FormItem className="grid grid-cols-4 items-center gap-4">
										<FormLabel className="text-right">Target Amount</FormLabel>
										<FormControl>
											<Input placeholder="1000" className="col-span-3" {...field} />
										</FormControl>
										{/* <FormDescription className="col-span-3 col-start-2">
											This is your public display name.
										</FormDescription> */}
										<FormMessage className="col-span-3 col-start-2" />
									</FormItem>
								)}
							/>
						</div>
						<DialogFooter>
							<Button type="submit">Submit</Button>
						</DialogFooter>
					</form>
				</Form>

			</DialogContent>
		</Dialog>
	)
}