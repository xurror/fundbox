import { fetcher, mapQueryParams, mutator } from "@/lib/api";
import { Contribution } from "@/types/contribution";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";

export function useContributions(params?: { fundId: string }) {
  const baseUri = "/api/contributions"
  const uri = params ? `${baseUri}?${mapQueryParams(params)}` : baseUri
  const {
    data,
    error,
    isLoading,
  } = useSWR<Array<Contribution>>(uri, fetcher)

  const { trigger: mutate, isMutating } = useSWRMutation(baseUri, mutator)

  return {
    data,
    error,
    isLoading,
    mutate,
    isMutating,
  }
}
