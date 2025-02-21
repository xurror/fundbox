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
import { Coins } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { toast } from "sonner"
import { getAccessToken } from "@auth0/nextjs-auth0"
import React from "react"

const FormSchema = z.object({
	fundId: z.string().uuid(),
	amount: z.coerce.number().gt(0, {
		message: 'Target amount must be at least 1.',
	}),
})

export function NewContributionForm({ fundId }: { fundId: string }) {
	const [open, setOpen] = React.useState(false);

	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: {
			fundId,
			amount: "" as any,
		},
	})

	async function onSubmit(data: z.infer<typeof FormSchema>) {
		console.log(data)
		const token = await getAccessToken();
		const response = await fetch('/api/contributions', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `Bearer ${token}`
			},
			body: JSON.stringify(data)
		})
		if (response.ok) {
			const contribution = await response.json()
			toast.success("Successfully made a contribution", {
				description: `You've successfully made a contribution of: ${contribution.amount}`,
			})
			form.reset()
			setOpen(false)
		} else {
			toast.error("Failed to make contribution", {
				description: "An error occurred while making your contribution.",
			})
		}
	}


	return (
		<Dialog open={open} onOpenChange={setOpen}>
			<DialogTrigger asChild>
				<Button className="ml-auto" >
					<Coins /> Make Contribution
				</Button>
			</DialogTrigger>
			<DialogContent className="sm:max-w-[512px]">
				<DialogHeader>
					<DialogTitle>Make a Contribution</DialogTitle>
					<DialogDescription>
						Let's make your contribution count.
					</DialogDescription>
				</DialogHeader>

				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)}>
						<div className="grid gap-4 py-4">
							<Input {...form.register("fundId", { required: true })} type="hidden" />

							<FormField
								control={form.control}
								name="amount"
								render={({ field }) => (
									<FormItem className="grid grid-cols-4 items-center gap-4">
										<FormLabel className="text-right">Amount</FormLabel>
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