import { fetcher } from "@/lib/api";
import useSWR from "swr";

type AppUser = {
  id: string
}

export function useAppUser() {
  const {
    data,
    error,
    isLoading,
  } = useSWR<AppUser>("/api/users/me", fetcher)

  return {
    data,
    error,
    isLoading,
  }
}
