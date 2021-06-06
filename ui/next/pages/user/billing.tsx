import React from 'react';
import { Layout } from 'components';
import PaymentMethod from 'components/Billing/PaymentMethod';
import Discounted from 'components/Billing/Discounted';

export function BillingContent() {
  return (
    <>
      <Discounted />

      <PaymentMethod />
    </>
  );
}

export default function Billing() {
  return <Layout content={BillingContent()} />;
}
