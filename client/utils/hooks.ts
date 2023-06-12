import { useState, useEffect } from "react";
import { TOKEN } from './constants';
import { useRouter } from "next/router";

export const useAuth = (settings?: {reroute: boolean, from?: string}) => {
  const router = useRouter();
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    validateToken()
  }, []);

  const validateToken = (refreshed?: boolean) => {
    const token = localStorage.getItem(TOKEN);
    //  TODO: could also check if token is expired
    if (token) setData(token)

    if (settings && settings.reroute && !token) {
      router.push(`/login${settings.from ? `?to=${settings.from}` : ''}`);
    }

    if (refreshed) setData(null)
  }

  const refresh = () => {
    validateToken(true)
  }

  return {token: data, refresh};
};

