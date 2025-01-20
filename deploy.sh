git pull origin master
cd ridnvil
npm install
npm run build
cd ..
docker compose build -n-no-cache ridnvil
docker compose up -d