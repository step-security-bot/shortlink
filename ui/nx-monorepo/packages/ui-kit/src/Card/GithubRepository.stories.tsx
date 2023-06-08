import { Meta } from '@storybook/react'

import GithubRepository from './GithubRepository'

const meta: Meta<any> = {
  title: 'UI/GithubRepository',
  component: GithubRepository,
}

export default meta

function Template(args) {
  return (
    <GithubRepository
      title="GitHub Repository"
      url="https://github.com/shortlink-org/shortlink"
      {...args}
    />
  )
}

export const Default = Template.bind({})
Default.args = {}
