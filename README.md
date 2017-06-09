# Not AWK (or sed, or perl)

Here are things that I would have done very quickly in awk or sed or perl or some combination of them. But I'm trying to teach myself how to use golang. Perhaps this effort will pay off in time saved eventually.

## csv2md

csv2md is to convert CSV files to markdown tables. It has no options or customizations, and the hard work is done by the encoding/csv package. But I'm trying to just write little filters and I had a need for this.