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
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { toast } from "sonner"
import React from "react"
import { useParams } from "next/navigation"
import { useFunds } from "@/hooks/use-funds"
import { useContributions } from "@/hooks/use-contributions"
import { Contribution } from "@/types/contribution"
import { EmbeddedCheckout, EmbeddedCheckoutProvider } from "@stripe/react-stripe-js"
import { loadStripe } from "@stripe/stripe-js"
import { fetcher } from "@/lib/api"


const FormSchema = z.object({
	fundId: z.string().nonempty("Please pick a fund to contribute to.").uuid(),
	amount: z.coerce.number().gt(0, {
		message: 'Target amount must be at least 1.',
	}),
})

export function NewContributionForm() {
	const { fundId } = useParams()
	const [stripeClientSecret, setStripeClientSecret] = React.useState<string | null>(null)
	const { data: funds } = useFunds()
	const { mutate } = useContributions()

	const [open, setOpen] = React.useState(false);

	React.useEffect(() => {
		fetcher<{ clientSecret: string }>("/api/stripe/checkout-session", {
			method: "POST",
			body: JSON.stringify({ account: "acct_1Qx7X8QOnFleJR7j" }),
		}).then((json) => {
			setStripeClientSecret(json.clientSecret)
		})
	}, [])

	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: {
			fundId: fundId as string,
			amount: "" as unknown as number,
		},
	})

	async function onSubmit(data: z.infer<typeof FormSchema>) {
		mutate(JSON.stringify(data)).then((contribution) => {
			toast.success("Successfully made a contribution", {
				description: `You've successfully made a contribution of: ${(contribution as Contribution).amount}`,
			})
			form.reset()
			setOpen(false)
		}).catch(() => {
			toast.error("Failed to make contribution", {
				description: "An error occurred while making your contribution.",
			})
		})
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
						Let&apos;s make your contribution count.
					</DialogDescription>
				</DialogHeader>
				{stripeClientSecret ?? (
					<EmbeddedCheckoutProvider
						stripe={loadStripe('STRIPE_PK', {
							stripeAccount: "STRIPE_ACCOUNT_KEY",
						})}
						options={{ clientSecret: stripeClientSecret }}
					>
						<EmbeddedCheckout />
					</EmbeddedCheckoutProvider>
				)}
				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)}>
						<div className="grid gap-4 py-4">
							<FormField
								disabled={!!fundId}
								control={form.control}
								name="fundId"
								render={({ field }) => (
									<FormItem className="grid grid-cols-4 items-center gap-4">
										<FormLabel className="text-right">Fund</FormLabel>
										<Select disabled={!!fundId} onValueChange={field.onChange} defaultValue={field.value}>
											<FormControl>
												<SelectTrigger className="col-span-3">
													<SelectValue placeholder="Select a fund to contribute to" />
												</SelectTrigger>
											</FormControl>
											<SelectContent>
												{funds?.map(fund => (
													<SelectItem key={fund.id} value={fund.id}>{fund.name}</SelectItem>
												))}
											</SelectContent>
										</Select>
										<FormMessage className="col-span-3 col-start-2" />
									</FormItem>
								)}
							/>

							<FormField
								control={form.control}
								name="amount"
								render={({ field }) => (
									<FormItem className="grid grid-cols-4 items-center gap-4">
										<FormLabel className="text-right">Amount</FormLabel>
										<FormControl>
											<Input placeholder="1000" className="col-span-3" {...field} />
										</FormControl>
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
