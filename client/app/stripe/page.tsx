"use client";

import React, { useState } from "react";
import {
  ConnectAccountOnboarding,
  ConnectComponentsProvider,
} from "@stripe/react-connect-js";
import useStripeConnect from "@/hooks/use-stripe-connect";
import { fetcher } from "@/lib/api";
import { Button } from "@/components/ui/button";

export default function Page() {
  const [accountLinkCreatePending, setAccountLinkCreatePending] = useState(false);
  const [accountCreatePending, setAccountCreatePending] = useState(false);
  const [onboardingExited, setOnboardingExited] = useState(false);
  const [error, setError] = useState(false);
  const [connectedAccountId, setConnectedAccountId] = useState<string>();
  const stripeConnectInstance = useStripeConnect(connectedAccountId);

  return (
    <div className="container">
      <div className="banner">
        <h2>Rocket Rides</h2>
      </div>
      <div className="content">
        {!connectedAccountId && <h2>Get ready for take off</h2>}
        {connectedAccountId && !stripeConnectInstance && <h2>Add information to start accepting money</h2>}
        {!connectedAccountId && <p>Rocket Rides is the world's leading air travel platform: join our team of pilots to help people travel faster.</p>}
        {!accountCreatePending && !connectedAccountId && (
          <div>
            <Button
              onClick={async () => {
                setAccountCreatePending(true);
                setError(false);
                fetcher("/api/stripe/account", {
                  method: "POST",
                })
                  .then((json) => {
                    setAccountCreatePending(false);
                    const { account, error } = json as any;

                    if (account) {
                      setConnectedAccountId(account);
                    }

                    if (error) {
                      setError(true);
                    }
                  });
              }}
            >
              Sign up
            </Button>
          </div>
        )}
        {connectedAccountId && !accountLinkCreatePending && (
          <Button
            onClick={async () => {
              setAccountLinkCreatePending(true);
              setError(false);
              fetcher("/api/stripe/account-link", {
                method: "POST",
                body: JSON.stringify({ account: connectedAccountId }),
              }).then((json) => {
                setAccountLinkCreatePending(false);
                const { url, error } = json as any;
                if (url) {
                  window.location.href = url;
                }

                if (error) {
                  setError(true);
                }
              });
            }}
          >
            Add information
          </Button>
        )}
        {stripeConnectInstance && (
          <ConnectComponentsProvider connectInstance={stripeConnectInstance}>
            <ConnectAccountOnboarding
              onExit={() => setOnboardingExited(true)}
            />
          </ConnectComponentsProvider>
        )}
        {error && <p className="error">Something went wrong!</p>}
        {(connectedAccountId || accountCreatePending || onboardingExited) && (
          <div className="dev-callout">
            {connectedAccountId && <p>Your connected account ID is: <code className="bold">{connectedAccountId}</code></p>}
            {accountCreatePending && <p>Creating a connected account...</p>}
            {onboardingExited && <p>The Account Onboarding component has exited</p>}
          </div>
        )}
        <div className="info-callout">
          <p>
            This is a sample app for Connect onboarding using the Account Onboarding embedded component. <a href="https://docs.stripe.com/connect/onboarding/quickstart?connect-onboarding-surface=embedded" target="_blank" rel="noopener noreferrer">View docs</a>
          </p>
        </div>
      </div>
    </div>
  );
}
