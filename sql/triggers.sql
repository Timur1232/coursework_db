CREATE OR REPLACE FUNCTION ensure_single_commander_per_vgk()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.position = 'commander' AND NEW.id_vgk IS NOT NULL THEN
    IF EXISTS (
      SELECT 1 FROM vgk_rescuers 
      WHERE id_vgk = NEW.id_vgk 
        AND position = 'commander' 
        AND id_rescuer != NEW.id_rescuer
    ) THEN
      RAISE EXCEPTION 'В одном ВГК может быть только один командир';
    END IF;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER single_commander_check
BEFORE INSERT OR UPDATE ON vgk_rescuers
FOR EACH ROW EXECUTE FUNCTION ensure_single_commander_per_vgk();
