# fantasy
An ETL pipeline and PWA for a consolidated Fantasy Football information repository.

Fantasy pulls the latest information from Walter Football (currently every hour). The information is pulled in Excel spreadsheet format, parsed to CSV for file storage, and transformed into Postgres.

The latest information can be viewed in a Progressive Web App (built in React/TypeScript) that users can quickly reference from their phone or computer.
