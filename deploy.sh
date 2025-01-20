git pull origin master
npm install
npm run build
docker compose build -n-no-cache ridnvil
docker compose up -d