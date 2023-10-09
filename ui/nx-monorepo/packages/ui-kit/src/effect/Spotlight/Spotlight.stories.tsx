import { Meta } from '@storybook/react'

import Spotlight, { SpotlightCard } from './Spotlight'

const meta: Meta<any> = {
  title: 'Effect/Spotlight',
  component: Spotlight,
}

export default meta

function Template(args: any) {
  return (
    <Spotlight className="max-w-sm mx-auto grid gap-6 lg:grid-cols-3 items-start lg:max-w-none group">
      <SpotlightCard>
        <div className="relative h-full bg-slate-900 p-6 pb-8 rounded-[inherit] z-20 overflow-hidden">
          {/* Radial gradient */}
          <div
            className="absolute bottom-0 translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/2 aspect-square"
            aria-hidden="true"
          >
            <div className="absolute inset-0 translate-z-0 bg-slate-800 rounded-full blur-[80px]" />
          </div>
          <div className="flex flex-col h-full items-center text-center">
            {/* Image */}
            <div className="relative inline-flex">
              <div
                className="w-[40%] h-[40%] absolute inset-0 m-auto -translate-y-[10%] blur-3xl -z-10 rounded-full bg-indigo-600"
                aria-hidden="true"
              />
            </div>
            {/* Text */}
            <div className="grow mb-5">
              <h2 className="text-xl text-slate-200 font-bold mb-1">
                Amazing Integration
              </h2>
              <p className="text-sm text-slate-500">
                Quickly apply filters to refine your issues lists and create
                custom views.
              </p>
            </div>
            <a
              className="inline-flex justify-center items-center whitespace-nowrap rounded-lg bg-slate-800 hover:bg-slate-900 border border-slate-700 px-3 py-1.5 text-sm font-medium text-slate-300 focus-visible:outline-none focus-visible:ring focus-visible:ring-indigo-300 dark:focus-visible:ring-slate-600 transition-colors duration-150"
              href="#0"
            >
              <svg
                className="fill-slate-500 mr-2"
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="14"
              >
                <path d="M12.82 8.116A.5.5 0 0 0 12 8.5V10h-.185a3 3 0 0 1-2.258-1.025l-.4-.457-1.328 1.519.223.255A5 5 0 0 0 11.815 12H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM12.82.116A.5.5 0 0 0 12 .5V2h-.185a5 5 0 0 0-3.763 1.708L3.443 8.975A3 3 0 0 1 1.185 10H1a1 1 0 1 0 0 2h.185a5 5 0 0 0 3.763-1.708l4.609-5.267A3 3 0 0 1 11.815 4H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM1 4h.185a3 3 0 0 1 2.258 1.025l.4.457 1.328-1.52-.223-.254A5 5 0 0 0 1.185 2H1a1 1 0 0 0 0 2Z" />
              </svg>
              <span>Connect</span>
            </a>
          </div>
        </div>
      </SpotlightCard>

      <SpotlightCard>
        <div className="relative h-full bg-slate-900 p-6 pb-8 rounded-[inherit] z-20 overflow-hidden">
          {/* Radial gradient */}
          <div
            className="absolute bottom-0 translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/2 aspect-square"
            aria-hidden="true"
          >
            <div className="absolute inset-0 translate-z-0 bg-slate-800 rounded-full blur-[80px]" />
          </div>
          <div className="flex flex-col h-full items-center text-center">
            {/* Image */}
            <div className="relative inline-flex">
              <div
                className="w-[40%] h-[40%] absolute inset-0 m-auto -translate-y-[10%] blur-3xl -z-10 rounded-full bg-indigo-600"
                aria-hidden="true"
              />
            </div>
            {/* Text */}
            <div className="grow mb-5">
              <h2 className="text-xl text-slate-200 font-bold mb-1">
                Amazing Integration
              </h2>
              <p className="text-sm text-slate-500">
                Quickly apply filters to refine your issues lists and create
                custom views.
              </p>
            </div>
            <a
              className="inline-flex justify-center items-center whitespace-nowrap rounded-lg bg-slate-800 hover:bg-slate-900 border border-slate-700 px-3 py-1.5 text-sm font-medium text-slate-300 focus-visible:outline-none focus-visible:ring focus-visible:ring-indigo-300 dark:focus-visible:ring-slate-600 transition-colors duration-150"
              href="#0"
            >
              <svg
                className="fill-slate-500 mr-2"
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="14"
              >
                <path d="M12.82 8.116A.5.5 0 0 0 12 8.5V10h-.185a3 3 0 0 1-2.258-1.025l-.4-.457-1.328 1.519.223.255A5 5 0 0 0 11.815 12H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM12.82.116A.5.5 0 0 0 12 .5V2h-.185a5 5 0 0 0-3.763 1.708L3.443 8.975A3 3 0 0 1 1.185 10H1a1 1 0 1 0 0 2h.185a5 5 0 0 0 3.763-1.708l4.609-5.267A3 3 0 0 1 11.815 4H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM1 4h.185a3 3 0 0 1 2.258 1.025l.4.457 1.328-1.52-.223-.254A5 5 0 0 0 1.185 2H1a1 1 0 0 0 0 2Z" />
              </svg>
              <span>Connect</span>
            </a>
          </div>
        </div>
      </SpotlightCard>

      <SpotlightCard>
        <div className="relative h-full bg-slate-900 p-6 pb-8 rounded-[inherit] z-20 overflow-hidden">
          {/* Radial gradient */}
          <div
            className="absolute bottom-0 translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-1/2 aspect-square"
            aria-hidden="true"
          >
            <div className="absolute inset-0 translate-z-0 bg-slate-800 rounded-full blur-[80px]" />
          </div>
          <div className="flex flex-col h-full items-center text-center">
            {/* Image */}
            <div className="relative inline-flex">
              <div
                className="w-[40%] h-[40%] absolute inset-0 m-auto -translate-y-[10%] blur-3xl -z-10 rounded-full bg-indigo-600"
                aria-hidden="true"
              />
            </div>
            {/* Text */}
            <div className="grow mb-5">
              <h2 className="text-xl text-slate-200 font-bold mb-1">
                Amazing Integration
              </h2>
              <p className="text-sm text-slate-500">
                Quickly apply filters to refine your issues lists and create
                custom views.
              </p>
            </div>
            <a
              className="inline-flex justify-center items-center whitespace-nowrap rounded-lg bg-slate-800 hover:bg-slate-900 border border-slate-700 px-3 py-1.5 text-sm font-medium text-slate-300 focus-visible:outline-none focus-visible:ring focus-visible:ring-indigo-300 dark:focus-visible:ring-slate-600 transition-colors duration-150"
              href="#0"
            >
              <svg
                className="fill-slate-500 mr-2"
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="14"
              >
                <path d="M12.82 8.116A.5.5 0 0 0 12 8.5V10h-.185a3 3 0 0 1-2.258-1.025l-.4-.457-1.328 1.519.223.255A5 5 0 0 0 11.815 12H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM12.82.116A.5.5 0 0 0 12 .5V2h-.185a5 5 0 0 0-3.763 1.708L3.443 8.975A3 3 0 0 1 1.185 10H1a1 1 0 1 0 0 2h.185a5 5 0 0 0 3.763-1.708l4.609-5.267A3 3 0 0 1 11.815 4H12v1.5a.5.5 0 0 0 .82.384l3-2.5a.5.5 0 0 0 0-.768l-3-2.5ZM1 4h.185a3 3 0 0 1 2.258 1.025l.4.457 1.328-1.52-.223-.254A5 5 0 0 0 1.185 2H1a1 1 0 0 0 0 2Z" />
              </svg>
              <span>Connect</span>
            </a>
          </div>
        </div>
      </SpotlightCard>
    </Spotlight>
  )
}

export const Default = Template.bind({})

// @ts-ignore
Default.args = { className: 'p-4 bg-gray-200' }
