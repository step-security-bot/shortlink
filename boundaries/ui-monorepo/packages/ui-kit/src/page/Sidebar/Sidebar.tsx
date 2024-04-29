import * as React from 'react'
import CollapsibleMenu from './CollapsibleMenu/CollapsibleMenu'

// Importing icons
import AddLinkIcon from '@mui/icons-material/AddLink'
import SettingsIcon from '@mui/icons-material/Settings'
import ReceiptIcon from '@mui/icons-material/Receipt'
import HttpIcon from '@mui/icons-material/Http'
import DashboardIcon from '@mui/icons-material/Dashboard'
import ListIcon from '@mui/icons-material/List'
import BarChartIcon from '@mui/icons-material/BarChart'
import PersonIcon from '@mui/icons-material/Person'
import SecurityIcon from '@mui/icons-material/Security'
import LayersIcon from '@mui/icons-material/Layers'
import AccountBalanceWalletIcon from '@mui/icons-material/AccountBalanceWallet'
import AssessmentIcon from '@mui/icons-material/Assessment'
import AssignmentIcon from '@mui/icons-material/Assignment'
import PeopleIcon from '@mui/icons-material/People'
import GroupAddIcon from '@mui/icons-material/GroupAdd'
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart'
import AdminPanelSettingsIcon from '@mui/icons-material/AdminPanelSettings'

type AppProps = {
  mode: 'full' | 'compact'
}

const iconClassName =
  'w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white'
const linkClassName =
  'flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group'

export function Sidebar({ mode }: AppProps) {
  return (
    <aside
      id="default-sidebar"
      className="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
      aria-label="Sidebar"
    >
      <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
        <ul className="space-y-2 font-medium">
          <li>
            <a href="/user/addUrl" className={linkClassName}>
              <AddLinkIcon className={iconClassName} />
              <span className="ms-3">Add URL</span>
            </a>
          </li>

          <CollapsibleMenu icon={ListIcon} title="Links">
            <li>
              <a href="/user/dashboard" className={linkClassName}>
                <DashboardIcon className={iconClassName} />
                <span className="ms-3">Dashboard</span>
              </a>
            </li>

            <li>
              <a href="/user/links" className={linkClassName}>
                <HttpIcon className={iconClassName} />
                <span className="ms-3">Links</span>
              </a>
            </li>

            <li>
              <a href="/user/reports" className={linkClassName}>
                <BarChartIcon className={iconClassName} />
                <span className="ms-3">Reports</span>
              </a>
            </li>
          </CollapsibleMenu>

          <CollapsibleMenu icon={SettingsIcon} title="Setting">
            <li>
              <a href="/user/profile" className={linkClassName}>
                <PersonIcon className={iconClassName} />
                <span className="ms-3">Profile</span>
              </a>
            </li>

            <li>
              <a href="/user/security" className={linkClassName}>
                <SecurityIcon className={iconClassName} />
                <span className="ms-3">Security</span>
              </a>
            </li>

            <li>
              <a href="/user/integrations" className={linkClassName}>
                <LayersIcon className={iconClassName} />
                <span className="ms-3">Integrations</span>
              </a>
            </li>

            <li>
              <a href="/user/audit" className={linkClassName}>
                <AssessmentIcon className={iconClassName} />
                <span className="ms-3">Audit</span>
              </a>
            </li>
          </CollapsibleMenu>

          <CollapsibleMenu icon={ShoppingCartIcon} title="E-commerce">
            <li>
              <a href="/user/billing" className={linkClassName + ' m-2'}>
                <AccountBalanceWalletIcon className={iconClassName} />
                <span className="ms-3">Billing</span>
              </a>
            </li>

            <li>
              <a href="/user/invoice" className={linkClassName + ' m-2'}>
                <ReceiptIcon className={iconClassName} />
                <span className="ms-3">Invoice</span>
              </a>
            </li>
          </CollapsibleMenu>

          <CollapsibleMenu icon={AdminPanelSettingsIcon} title="Admin">
            <li>
              <a href="/admin/links" className={linkClassName}>
                <ListIcon className={iconClassName} />
                <span className="ms-3">Links</span>
              </a>
            </li>

            <li className={'border-t border-gray-200 dark:border-gray-70'}>
              <a href="/admin/users" className={linkClassName}>
                <GroupAddIcon className={iconClassName} />
                <span className="ms-3">Users</span>
              </a>
            </li>

            <li className={'border-b border-gray-200 dark:border-gray-70'}>
              <a href="/admin/groups" className={linkClassName}>
                <PeopleIcon className={iconClassName} />
                <span className="ms-3">Groups</span>
              </a>
            </li>
          </CollapsibleMenu>

          <li>
            <a href="/about" className={linkClassName}>
              <AssignmentIcon className={iconClassName} />
              <span className="ms-3">About As</span>
            </a>
          </li>
        </ul>
      </div>
    </aside>
  )
}

export default Sidebar
