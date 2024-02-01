import dotenv from 'dotenv'

dotenv.config()


if (!process.env.BASE_URL) {
  throw new Error('BASE_URL is not defined')
}

export const projectBaseUrl = process.env.BASE_URL!
