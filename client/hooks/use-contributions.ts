import { fetcher } from "@/lib/api";
import { Contribution } from "@/types/contribution";
import useSWR from "swr";

export function useContributions(fundId?: string) {
  const {
    data,
    error,
    isLoading
  } = useSWR<Array<Contribution>>(
    "/api/contributions" + (fundId ? `?fundId=${fundId}` : ""),
    fetcher
  )

  return {
    data,
    error,
    isLoading
  }
}