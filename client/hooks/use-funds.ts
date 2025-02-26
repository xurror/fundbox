import { fetcher, mapQueryParams, mutator } from "@/lib/api";
import { Fund } from "@/types/fund";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";

export function useFunds(params?: { contributorId: string }) {
  const baseUri = "/api/funds"
  const uri = params ? `${baseUri}?${mapQueryParams(params)}` : baseUri

  const {
    data,
    error,
    isLoading,
  } = useSWR<Array<Fund>>(uri, fetcher)

  const { trigger: mutate } = useSWRMutation(baseUri, mutator)

  return {
    data,
    error,
    isLoading,
    mutate
  }
}

export function useFund(params: { id: string }) {
  const uri = `/api/funds/${params.id}`

  const {
    data,
    error,
    isLoading,
  } = useSWR<Fund>(uri, fetcher)

  return {
    data,
    error,
    isLoading,
  }
}
