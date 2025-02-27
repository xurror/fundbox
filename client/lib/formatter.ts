export const formatAmount = (amount: string | number) => {
  const amountFloat = typeof amount == 'string' ? parseFloat(amount) : amount
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
  }).format(amountFloat)
}

// 2025-02-21T19:41:54.174548+01:00
export const formatDate = (date: string) => {
  console.log(date)
  const dateObj = new Date(date)
  return dateObj.toLocaleDateString("en-US", {
    year: "numeric",
    month: "short",
    day: "numeric",
  })
}
