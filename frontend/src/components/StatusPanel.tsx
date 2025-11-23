import { CheckCircle2, Clock, ListChecks, PackageSearch } from 'lucide-react'
import type { ReactNode } from 'react'

import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

type StatusPanelProps = {
  isLoading?: boolean
  stats: {
    ingredients: number
    products: number
    soldOut: number
  }
  lastUpdated?: string | null
}

const StatusPanel = ({ isLoading, stats, lastUpdated }: StatusPanelProps) => {
  return (
    <Card className="h-full">
      <CardHeader className="gap-4">
        <div className="flex items-center gap-2">
          <ListChecks />
          <CardTitle className="text-xl">Match status</CardTitle>
        </div>
        <p className="text-sm text-foreground/70">
          We route the recipe through Gemini, parse ingredients, then run parallel Kosik searches.
        </p>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="grid gap-3">
          <StatusRow
            icon={<CheckCircle2 className="text-green-500" />}
            label="Ingredients parsed"
            value={stats.ingredients}
            hint="LLM extracted"
          />
          <StatusRow
            icon={<PackageSearch className="text-blue-500" />}
            label="Products surfaced"
            value={stats.products}
            hint="Kosik candidates"
          />
          <StatusRow
            icon={<Clock className="text-orange-500" />}
            label="Out of stock"
            value={stats.soldOut}
            hint="Flagged results"
          />
        </div>

        <div className="rounded-[var(--radius-base)] border-2 border-border bg-background px-4 py-3 shadow-[var(--shadow)]">
          <p className="text-xs uppercase text-foreground/70">Pipeline</p>
          <div className="mt-2 flex flex-wrap gap-2">
            <Badge variant="outline">Text → Gemini</Badge>
            <Badge variant="outline">Ingredient graph</Badge>
            <Badge variant="outline">Kosik fan-out</Badge>
          </div>
          <p className="mt-3 text-xs text-foreground/60">
            {isLoading ? 'Processing your recipe...' : lastUpdated ?? 'Waiting for input'}
          </p>
        </div>
      </CardContent>
    </Card>
  )
}

type StatusRowProps = {
  icon: ReactNode
  label: string
  value: number
  hint: string
}

const StatusRow = ({ icon, label, value, hint }: StatusRowProps) => (
  <div className="flex items-center justify-between rounded-[var(--radius-base)] border-2 border-border bg-background px-3 py-2 shadow-[var(--shadow)]">
    <div className="flex items-center gap-3">
      <div className="rounded-full border-2 border-border bg-secondary-background p-2">{icon}</div>
      <div>
        <p className="text-sm font-heading uppercase">{label}</p>
        <p className="text-xs text-foreground/70">{hint}</p>
      </div>
    </div>
    <span className="text-2xl font-heading">{value}</span>
  </div>
)

export default StatusPanel

