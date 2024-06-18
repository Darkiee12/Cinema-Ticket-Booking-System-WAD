import Auditorium from './auditorium';
import Cinema from "./cinema";
import Movie from './movie';

export default interface Show {
  id: string;
  auditorium: Auditorium;
  date: string;
  endTime: string;
  imdbID: string;
  startTime: string;
  movie: Movie;
}

export interface AuditoriumShow extends Auditorium {
  shows: Show[];
}
export interface CinemaAuditorium extends Cinema {
  auditoriums: AuditoriumShow[];
}
export class MovieShow {
  private cinemas: Map<string, CinemaAuditorium> = new Map();

  constructor(shows: Show[]) {
    for (const show of shows) {
      const cinemaId = show.auditorium.cinema.id;
      const auditoriumId = show.auditorium.id;

      // Create or get CinemaAuditorium from map
      let cinemaAuditorium = this.cinemas.get(cinemaId);
      if (!cinemaAuditorium) {
        cinemaAuditorium = {
          ...show.auditorium.cinema,
          auditoriums: [],
        };
        this.cinemas.set(cinemaId, cinemaAuditorium);
      }

      // Find or create AuditoriumShow
      let auditoriumShow = cinemaAuditorium.auditoriums.find(
        (aud) => aud.id === auditoriumId,
      );
      if (!auditoriumShow) {
        auditoriumShow = {
          ...show.auditorium,
          shows: [],
        };
        cinemaAuditorium.auditoriums.push(auditoriumShow);
      }

      // Add show to AuditoriumShow
      auditoriumShow.shows.push(show);
    }
  }

  getCinemas(date?: string): CinemaAuditorium[] {
    if (date) {
      return [...this.cinemas.values()]
        .map((cinema) => {
          cinema.auditoriums = cinema.auditoriums
            .map((auditorium) => {
              auditorium.shows = auditorium.shows.filter(
                (show) => show.date === date,
              );
              return auditorium;
            })
            .filter((auditorium) => auditorium.shows.length > 0);
          return cinema;
        })
        .filter((cinema) => cinema.auditoriums.length > 0);
    }
    return [...this.cinemas.values()];
  }
}
