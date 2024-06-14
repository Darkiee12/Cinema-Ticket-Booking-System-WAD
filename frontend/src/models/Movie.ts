export default interface Movie {
  imdbID: string
  awards: string
  dvd: string
  imdbRating: number
  imdbVotes: number
  metascore: number
  originalTitle: string
  plot: string
  poster: string
  rated: string
  released: string,
  runtime: number
  tagline: string
  title: string
  type: string
  website: string
  year: number
  production: string
  tmdbID: number
  boxOffice: number
  genres: Genre[]
  created_at: null
}

interface Genre{
  id: number
  name: string
}