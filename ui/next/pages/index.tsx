import React from 'react'
import type { NextPage } from 'next'
import Feature from 'components/Landing/feature'
import Header from 'components/Landing/header'
import Mobile from 'components/Landing/mobile'
import Subscribe from 'components/Landing/subscribe'
import Testimonials from 'components/Testimonials'

import { Layout } from 'components'

const Home: NextPage = () => {
  return (
    <Layout>
      <Header />
      <Mobile />
      <Feature />
      <Testimonials />
      <Subscribe />
    </Layout>
  )
}

export default Home;
